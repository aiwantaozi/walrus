package terraform

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/hashicorp/terraform-config-inspect/tfconfig"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/kubernetes"
	"k8s.io/utils/pointer"

	"github.com/seal-io/seal/pkg/auths"
	"github.com/seal-io/seal/pkg/auths/session"
	revisionbus "github.com/seal-io/seal/pkg/bus/servicerevision"
	"github.com/seal-io/seal/pkg/dao"
	"github.com/seal-io/seal/pkg/dao/model"
	"github.com/seal-io/seal/pkg/dao/model/connector"
	"github.com/seal-io/seal/pkg/dao/model/environmentconnectorrelationship"
	"github.com/seal-io/seal/pkg/dao/model/predicate"
	"github.com/seal-io/seal/pkg/dao/model/secret"
	"github.com/seal-io/seal/pkg/dao/model/service"
	"github.com/seal-io/seal/pkg/dao/model/serviceresource"
	"github.com/seal-io/seal/pkg/dao/model/servicerevision"
	"github.com/seal-io/seal/pkg/dao/model/templateversion"
	"github.com/seal-io/seal/pkg/dao/types"
	"github.com/seal-io/seal/pkg/dao/types/crypto"
	"github.com/seal-io/seal/pkg/dao/types/oid"
	"github.com/seal-io/seal/pkg/dao/types/property"
	"github.com/seal-io/seal/pkg/dao/types/status"
	deptypes "github.com/seal-io/seal/pkg/deployer/types"
	opk8s "github.com/seal-io/seal/pkg/operator/k8s"
	"github.com/seal-io/seal/pkg/settings"
	"github.com/seal-io/seal/pkg/terraform/config"
	"github.com/seal-io/seal/pkg/terraform/parser"
	"github.com/seal-io/seal/pkg/terraform/util"
	"github.com/seal-io/seal/utils/log"
)

// DeployerType the type of deployer.
const DeployerType = types.DeployerTypeTF

// Deployer terraform deployer to deploy the service.
type Deployer struct {
	logger      log.Logger
	modelClient model.ClientSet
	clientSet   *kubernetes.Clientset
}

type CreateRevisionOptions struct {
	// JobType indicates the type of the job.
	JobType string
	Tags    []string
	Service *model.Service
}

// CreateSecretsOptions options for creating deployment job secrets.
type CreateSecretsOptions struct {
	SkipTLSVerify   bool
	ServiceRevision *model.ServiceRevision
	Connectors      model.Connectors
	ProjectID       oid.ID
	// Metadata.
	ProjectName     string
	EnvironmentName string
	ServiceName     string
}

// CreateJobOptions options for do job action.
type CreateJobOptions struct {
	Type          string
	SkipTLSVerify bool
	Service       *model.Service
	// ServiceRevision indicates the service revision to create the deploy job.
	ServiceRevision *model.ServiceRevision
}

// _backendAPI the API path to terraform deploy backend.
// Terraform will get and update deployment states from this API.
const _backendAPI = "/v1/service-revisions/%s/terraform-states?projectID=%s"

// _secretPrefix the prefix of the secret name.
const _secretPrefix = "_seal_secret_"

// _servicePrefix the prefix of the service output name.
const _servicePrefix = "_seal_service_"

var (
	// _secretReg the regexp to match the secret variable.
	_secretReg = regexp.MustCompile(`\${secret\.([a-zA-Z0-9_]+)}`)
	// _serviceReg the regexp to match the service output.
	_serviceReg = regexp.MustCompile(`\${service\.([^.}]+)\.([^.}]+)}`)
)

func NewDeployer(_ context.Context, opts deptypes.CreateOptions) (deptypes.Deployer, error) {
	clientSet, err := kubernetes.NewForConfig(opts.KubeConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create kubernetes client set: %w", err)
	}

	return &Deployer{
		modelClient: opts.ModelClient,
		clientSet:   clientSet,
		logger:      log.WithName("deployer").WithName("tf"),
	}, nil
}

func (d Deployer) Type() deptypes.Type {
	return DeployerType
}

// Apply deploys the service.
func (d Deployer) Apply(ctx context.Context, service *model.Service, opts deptypes.ApplyOptions) (err error) {
	revision, err := d.CreateServiceRevision(ctx, CreateRevisionOptions{
		JobType: JobTypeApply,
		Tags:    opts.Tags,
		Service: service,
	})
	if err != nil {
		return err
	}

	defer func() {
		if err == nil {
			return
		}
		// Report to service revision.
		_ = d.updateRevisionStatus(ctx, revision, status.ServiceRevisionStatusFailed, err.Error())
	}()

	return d.CreateK8sJob(ctx, CreateJobOptions{
		Type:            JobTypeApply,
		SkipTLSVerify:   opts.SkipTLSVerify,
		Service:         service,
		ServiceRevision: revision,
	})
}

// Destroy will destroy the resource of the service.
// 1. Get the latest revision, and checkAppRevision it if it is running.
// 2. If not running, then destroy resources.
func (d Deployer) Destroy(
	ctx context.Context,
	service *model.Service,
	destroyOpts deptypes.DestroyOptions,
) (err error) {
	sr, err := d.CreateServiceRevision(ctx, CreateRevisionOptions{
		JobType: JobTypeDestroy,
		Service: service,
	})
	if err != nil {
		return err
	}

	defer func() {
		if err == nil {
			return
		}
		// Report to service revision.
		_ = d.updateRevisionStatus(ctx, sr, status.ServiceRevisionStatusFailed, err.Error())
	}()

	// If no resource exists, skip job and set revision status succeed.
	exist, err := d.modelClient.ServiceResources().Query().
		Where(serviceresource.ServiceID(service.ID)).
		Exist(ctx)
	if err != nil {
		return err
	}

	if !exist {
		return d.updateRevisionStatus(ctx, sr, status.ServiceRevisionStatusSucceeded, sr.StatusMessage)
	}

	return d.CreateK8sJob(ctx, CreateJobOptions{
		Type:            JobTypeDestroy,
		SkipTLSVerify:   destroyOpts.SkipTLSVerify,
		Service:         service,
		ServiceRevision: sr,
	})
}

// Rollback service to a specific revision.
func (d Deployer) Rollback(
	ctx context.Context,
	service *model.Service,
	opts deptypes.RollbackOptions,
) (err error) {
	if opts.CloneFrom == nil || opts.CloneFrom.ServiceID != service.ID {
		return errors.New("rollback failed: invalid revision")
	}

	status.ServiceStatusDeployed.Reset(service, "Rolling back")

	update, err := dao.ServiceUpdate(d.modelClient, service)
	if err != nil {
		return err
	}

	service, err = update.Save(ctx)
	if err != nil {
		return err
	}

	sr, err := d.CreateServiceRevision(ctx,
		CreateRevisionOptions{
			JobType: JobTypeApply,
			Tags:    opts.Tags,
			Service: service,
		},
	)
	if err != nil {
		return err
	}

	defer func() {
		if err == nil {
			return
		}

		if sr != nil {
			// Report to service revision.
			_ = d.updateRevisionStatus(ctx, sr, status.ServiceRevisionStatusFailed, err.Error())
			return
		}

		status.ServiceStatusDeployed.False(service, err.Error())

		serviceUpdate, updateErr := dao.ServiceUpdate(d.modelClient, service)
		if updateErr != nil {
			d.logger.Error(err)
			return
		}

		updateErr = serviceUpdate.Exec(ctx)
		if updateErr != nil {
			d.logger.Errorf("update service status failed: %v", updateErr)
		}
	}()

	return d.CreateK8sJob(ctx, CreateJobOptions{
		Type:            JobTypeApply,
		SkipTLSVerify:   opts.SkipTLSVerify,
		Service:         service,
		ServiceRevision: sr,
	})
}

// CreateK8sJob will create a k8s job to deploy、destroy or rollback the service.
func (d Deployer) CreateK8sJob(ctx context.Context, opts CreateJobOptions) error {
	connectors, err := d.getConnectors(ctx, opts.Service)
	if err != nil {
		return err
	}

	project, err := d.modelClient.Projects().Get(ctx, opts.Service.ProjectID)
	if err != nil {
		return err
	}

	environment, err := d.modelClient.Environments().Get(ctx, opts.Service.EnvironmentID)
	if err != nil {
		return err
	}

	// Prepare tfConfig for deployment.
	secretOpts := CreateSecretsOptions{
		SkipTLSVerify:   opts.SkipTLSVerify,
		ServiceRevision: opts.ServiceRevision,
		Connectors:      connectors,
		ProjectID:       opts.Service.ProjectID,
		// Metadata.
		ProjectName:     project.Name,
		EnvironmentName: environment.Name,
		ServiceName:     opts.Service.Name,
	}
	if err = d.createK8sSecrets(ctx, secretOpts); err != nil {
		return err
	}

	jobImage, err := settings.TerraformDeployerImage.Value(ctx, d.modelClient)
	if err != nil {
		return err
	}

	// Create deployment job.
	jobOpts := JobCreateOptions{
		Type:              opts.Type,
		ServiceRevisionID: opts.ServiceRevision.ID.String(),
		Image:             jobImage,
	}

	return CreateJob(ctx, d.clientSet, jobOpts)
}

func (d Deployer) updateRevisionStatus(ctx context.Context, ar *model.ServiceRevision, s, m string) error {
	// Report to service revision.
	ar.Status = s
	ar.StatusMessage = m

	update, err := dao.ServiceRevisionUpdate(d.modelClient, ar)
	if err != nil {
		return err
	}

	ar, err = update.Save(ctx)
	if err != nil {
		d.logger.Error(err)
		return err
	}

	if err = revisionbus.Notify(ctx, d.modelClient, ar); err != nil {
		d.logger.Error(err)
		return err
	}

	return nil
}

// createK8sSecrets will create the k8s secrets for deployment.
func (d Deployer) createK8sSecrets(ctx context.Context, opts CreateSecretsOptions) error {
	secretData := make(map[string][]byte)
	// SecretName terraform tfConfig name.
	secretName := _jobSecretPrefix + string(opts.ServiceRevision.ID)

	// Prepare terraform config files bytes for deployment.
	terraformData, err := d.LoadConfigsBytes(ctx, opts)
	if err != nil {
		return err
	}

	for k, v := range terraformData {
		secretData[k] = v
	}

	// Mount the provider configs(e.g. kubeconfig) to secret.
	providerData, err := d.GetProviderSecretData(opts.Connectors)
	if err != nil {
		return err
	}

	for k, v := range providerData {
		secretData[k] = v
	}

	// Create deployment secret.
	if err = CreateSecret(ctx, d.clientSet, secretName, secretData); err != nil {
		return err
	}

	return nil
}

// CreateServiceRevision will create a new service revision.
// Get the latest revision, and check it if it is running.
// If not running, then apply the latest revision.
// If running, then wait for the latest revision to be applied.
func (d Deployer) CreateServiceRevision(
	ctx context.Context,
	opts CreateRevisionOptions,
) (*model.ServiceRevision, error) {
	entity := &model.ServiceRevision{
		ProjectID:       opts.Service.ProjectID,
		ServiceID:       opts.Service.ID,
		EnvironmentID:   opts.Service.EnvironmentID,
		TemplateID:      opts.Service.Template.ID,
		TemplateVersion: opts.Service.Template.Version,
		Attributes:      opts.Service.Attributes,
		Tags:            opts.Tags,
		DeployerType:    DeployerType,
		Status:          status.ServiceRevisionStatusRunning,
	}

	// Output of the previous revision should be inherited to the new one
	// when creating a new revision.
	prevEntity, err := d.modelClient.ServiceRevisions().Query().
		Where(servicerevision.And(
			servicerevision.ServiceID(opts.Service.ID),
			servicerevision.DeployerType(DeployerType))).
		Order(model.Desc(servicerevision.FieldCreateTime)).
		First(ctx)
	if err != nil && !model.IsNotFound(err) {
		return nil, err
	}

	if prevEntity != nil {
		if prevEntity.Status == status.ServiceRevisionStatusRunning {
			return nil, errors.New("service deployment is running")
		}
		// Inherit the output of previous revision.
		entity.Output = prevEntity.Output

		requiredProviders, err := d.getRequiredProviders(ctx, opts.Service.ID, entity.Output)
		if err != nil {
			return nil, err
		}
		entity.PreviousRequiredProviders = requiredProviders
	}

	if opts.JobType == JobTypeDestroy &&
		prevEntity != nil &&
		prevEntity.Status == status.ServiceRevisionStatusSucceeded {
		entity.TemplateID = prevEntity.TemplateID
		entity.TemplateVersion = prevEntity.TemplateVersion
		entity.Attributes = prevEntity.Attributes
		entity.InputPlan = prevEntity.InputPlan
		entity.Output = prevEntity.Output
		entity.PreviousRequiredProviders = prevEntity.PreviousRequiredProviders
	}

	// Create revision, mark status to running.
	creates, err := dao.ServiceRevisionCreates(d.modelClient, entity)
	if err != nil {
		return nil, err
	}

	revision, err := creates[0].Save(ctx)
	if err != nil {
		return nil, err
	}
	revision.Edges.Service = opts.Service

	return revision, nil
}

func (d Deployer) getRequiredProviders(
	ctx context.Context,
	instanceID oid.ID,
	previousOutput string,
) ([]types.ProviderRequirement, error) {
	stateRequiredProviderSet := sets.NewString()

	previousRequiredProviders, err := d.getPreviousRequiredProviders(ctx, instanceID)
	if err != nil {
		return nil, err
	}

	stateRequiredProviders, err := parser.ParseStateProviders(previousOutput)
	if err != nil {
		return nil, err
	}

	stateRequiredProviderSet.Insert(stateRequiredProviders...)

	requiredProviders := make([]types.ProviderRequirement, 0, len(previousRequiredProviders))

	for _, p := range previousRequiredProviders {
		if stateRequiredProviderSet.Has(p.Name) {
			requiredProviders = append(requiredProviders, p)
		}
	}

	return requiredProviders, nil
}

// LoadConfigsBytes returns terraform main.tf and terraform.tfvars for deployment.
func (d Deployer) LoadConfigsBytes(ctx context.Context, opts CreateSecretsOptions) (map[string][]byte, error) {
	logger := log.WithName("deployer").WithName("tf")
	// Prepare terraform tfConfig.
	//  get module configs from service revision.
	moduleConfigs, providerRequirements, err := d.GetModuleConfigs(ctx, opts)
	if err != nil {
		return nil, err
	}
	// Merge current and previous required providers.
	providerRequirements = append(providerRequirements,
		opts.ServiceRevision.PreviousRequiredProviders...)

	requiredProviders := make(map[string]*tfconfig.ProviderRequirement, 0)
	for _, p := range providerRequirements {
		if _, ok := requiredProviders[p.Name]; !ok {
			requiredProviders[p.Name] = p.ProviderRequirement
		} else {
			logger.Warnf("duplicate provider requirement: %s", p.Name)
		}
	}

	// Parse module attributes.
	secrets, dependencyOutputs, err := d.parseModuleAttributes(ctx, moduleConfigs, opts)
	if err != nil {
		return nil, err
	}

	// Prepare address for terraform backend.
	serverAddress, err := settings.ServeUrl.Value(ctx, d.modelClient)
	if err != nil {
		return nil, err
	}

	if serverAddress == "" {
		return nil, errors.New("server address is empty")
	}
	address := fmt.Sprintf("%s%s", serverAddress,
		fmt.Sprintf(_backendAPI, opts.ServiceRevision.ID, opts.ProjectID))

	// Prepare API token for terraform backend.
	const _30mins = 1800

	at, err := auths.CreateAccessToken(ctx,
		d.modelClient, types.TokenKindDeployment, string(opts.ServiceRevision.ID), pointer.Int(_30mins))
	if err != nil {
		return nil, err
	}

	// Prepare terraform config files to be mounted to secret.
	requiredProviderNames := sets.NewString()
	for _, p := range providerRequirements {
		requiredProviderNames = requiredProviderNames.Insert(p.Name)
	}

	secretNames := make([]string, 0, len(secrets))
	for _, s := range secrets {
		secretNames = append(secretNames, s.Name)
	}

	// Prepare outputs.
	var outputCount int
	for _, v := range moduleConfigs {
		outputCount += len(v.Outputs)
	}

	outputs := make([]config.Output, 0, outputCount)
	for _, v := range moduleConfigs {
		outputs = append(outputs, v.Outputs...)
	}
	secretOptionMaps := map[string]config.CreateOptions{
		config.FileMain: {
			TerraformOptions: &config.TerraformOptions{
				Token:                at.Value,
				Address:              address,
				SkipTLSVerify:        opts.SkipTLSVerify,
				ProviderRequirements: requiredProviders,
			},
			ProviderOptions: &config.ProviderOptions{
				RequiredProviderNames: requiredProviderNames.List(),
				Connectors:            opts.Connectors,
				SecretMonthPath:       _secretMountPath,
				ConnectorSeparator:    parser.ConnectorSeparator,
			},
			ModuleOptions: &config.ModuleOptions{
				ModuleConfigs: moduleConfigs,
			},
			VariableOptions: &config.VariableOptions{
				SecretPrefix:      _secretPrefix,
				ServicePrefix:     _servicePrefix,
				SecretNames:       secretNames,
				DependencyOutputs: dependencyOutputs,
			},
			OutputOptions: outputs,
		},
		config.FileVars: getVarConfigOptions(secrets, dependencyOutputs),
	}
	secretMaps := make(map[string][]byte, 0)

	for k, v := range secretOptionMaps {
		secretMaps[k], err = config.CreateConfigToBytes(v)
		if err != nil {
			return nil, err
		}
	}

	// Save input plan to service revision.
	opts.ServiceRevision.InputPlan = string(secretMaps[config.FileMain])
	// If service revision does not inherit secrets from cloned revision,
	// then save the parsed secrets to service revision.
	if len(opts.ServiceRevision.Secrets) == 0 {
		secretMap := make(crypto.Map[string, string], len(secrets))
		for _, s := range secrets {
			secretMap[s.Name] = string(s.Value)
		}
		opts.ServiceRevision.Secrets = secretMap
	}

	update, err := dao.ServiceRevisionUpdate(d.modelClient, opts.ServiceRevision)
	if err != nil {
		return nil, err
	}

	revision, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}

	if err = revisionbus.Notify(ctx, d.modelClient, revision); err != nil {
		return nil, err
	}

	return secretMaps, nil
}

// GetProviderSecretData returns provider kubeconfig secret data mount into terraform container.
func (d Deployer) GetProviderSecretData(connectors model.Connectors) (map[string][]byte, error) {
	secretData := make(map[string][]byte)

	for _, c := range connectors {
		if c.Type != types.ConnectorTypeK8s {
			continue
		}

		_, s, err := opk8s.LoadApiConfig(*c)
		if err != nil {
			return nil, err
		}

		// NB(alex) the secret file name must be config + connector id to
		// match with terraform provider in config convert.
		secretFileName := util.GetK8sSecretName(c.ID.String())
		secretData[secretFileName] = []byte(s)
	}

	return secretData, nil
}

// GetModuleConfigs returns module configs and required connectors to
// get terraform module config block from service revision.
func (d Deployer) GetModuleConfigs(
	ctx context.Context,
	opts CreateSecretsOptions,
) ([]*config.ModuleConfig, []types.ProviderRequirement, error) {
	var (
		requiredProviders = make([]types.ProviderRequirement, 0)
		predicates        = make([]predicate.TemplateVersion, 0)
	)

	predicates = append(predicates, templateversion.And(
		templateversion.TemplateID(opts.ServiceRevision.TemplateID),
		templateversion.Version(opts.ServiceRevision.TemplateVersion),
	))

	moduleVersion, err := d.modelClient.TemplateVersions().
		Query().
		Select(
			templateversion.FieldID,
			templateversion.FieldTemplateID,
			templateversion.FieldVersion,
			templateversion.FieldSource,
			templateversion.FieldSchema,
		).
		Where(templateversion.Or(predicates...)).
		Only(ctx)
	if err != nil {
		return nil, nil, err
	}

	if moduleVersion.Schema != nil {
		requiredProviders = append(requiredProviders, moduleVersion.Schema.RequiredProviders...)
	}

	mc, err := getModuleConfig(opts.ServiceRevision, moduleVersion, opts)
	if err != nil {
		return nil, nil, err
	}

	return []*config.ModuleConfig{mc}, requiredProviders, err
}

func (d Deployer) getConnectors(ctx context.Context, ai *model.Service) (model.Connectors, error) {
	rs, err := d.modelClient.EnvironmentConnectorRelationships().Query().
		Where(environmentconnectorrelationship.EnvironmentID(ai.EnvironmentID)).
		WithConnector(func(cq *model.ConnectorQuery) {
			cq.Select(
				connector.FieldID,
				connector.FieldName,
				connector.FieldType,
				connector.FieldCategory,
				connector.FieldConfigVersion,
				connector.FieldConfigData)
		}).
		All(ctx)
	if err != nil {
		return nil, err
	}

	var cs model.Connectors
	for i := range rs {
		cs = append(cs, rs[i].Edges.Connector)
	}

	return cs, nil
}

// parseModuleAttributes parse module secrets and dependencies, return matched model.Secrets and service output.
func (d Deployer) parseModuleAttributes(
	ctx context.Context,
	templateConfig []*config.ModuleConfig,
	opts CreateSecretsOptions,
) (secrets model.Secrets, outputs map[string]parser.OutputState, err error) {
	var (
		templateSecrets          []string
		dependencyServiceOutputs []string
	)

	for _, moduleConfig := range templateConfig {
		templateSecrets, dependencyServiceOutputs = parseAttributeReplace(
			moduleConfig.Attributes,
			templateSecrets,
			dependencyServiceOutputs,
		)
	}
	// If service revision has secrets that inherit from cloned revision, use them directly.
	if len(opts.ServiceRevision.Secrets) > 0 {
		for k, v := range opts.ServiceRevision.Secrets {
			secrets = append(secrets, &model.Secret{
				Name:  k,
				Value: crypto.String(v),
			})
		}
	} else {
		secrets, err = d.getSecrets(ctx, templateSecrets, opts.ProjectID)
		if err != nil {
			return nil, nil, err
		}
	}

	outputs, err = d.getServiceDependencyOutputs(ctx, opts.ServiceRevision.ServiceID, dependencyServiceOutputs)
	if err != nil {
		return nil, nil, err
	}

	return secrets, outputs, nil
}

func (d Deployer) getSecrets(ctx context.Context, secretNames []string, projectID oid.ID) (model.Secrets, error) {
	var secrets model.Secrets

	nameIn := make([]interface{}, len(secretNames))
	for i, name := range secretNames {
		nameIn[i] = name
	}
	// This query is used to distinct the secrets with the same name.
	//  SELECT
	//    "id",
	//    "name",
	//    "value"
	//  FROM
	//    "secrets"
	//  WHERE
	//    (
	//      (
	//        "project_id" IS NULL
	//        AND "name" NOT IN (
	//          SELECT
	//            "name"
	//          FROM
	//            "secrets"
	//          WHERE
	//            "project_id" = opts.ProjectID
	//        )
	//      )
	//      OR "project_id" = opts.ProjectID
	//    )
	//    AND NAME IN (moduleSecrets)
	err := func() error {
		s, err := session.GetSubject(ctx)
		if err == nil {
			s.IncognitoOn()
			defer s.IncognitoOff()
		}

		return d.modelClient.Secrets().Query().
			Modify(func(s *sql.Selector) {
				// Select secrets without project id or not in project.
				subQuery := sql.Select(secret.FieldName).
					From(sql.Table(secret.Table)).
					Where(sql.EQ(secret.FieldProjectID, projectID))
				s.Select(secret.FieldID, secret.FieldName, secret.FieldValue).
					Where(
						sql.And(
							sql.Or(
								sql.And(
									sql.IsNull(secret.FieldProjectID),
									sql.NotIn(secret.FieldName, subQuery),
								),
								sql.EQ(secret.FieldProjectID, projectID),
							),
							sql.In(secret.FieldName, nameIn...),
						),
					)
			}).
			Scan(ctx, &secrets)
	}()
	if err != nil {
		return nil, err
	}

	// Validate module secret are all exist.
	foundSecretSet := sets.NewString()
	for _, s := range secrets {
		foundSecretSet.Insert(s.Name)
	}
	requiredSecretSet := sets.NewString(secretNames...)

	missingSecretSet := requiredSecretSet.Difference(foundSecretSet)
	if missingSecretSet.Len() > 0 {
		return nil, fmt.Errorf("missing secrets: %s", missingSecretSet.List())
	}

	return secrets, nil
}

// getPreviousRequiredProviders get previous succeed revision required providers.
// NB(alex): the previous revision may be failed, the failed revision may not contain required providers of states.
func (d Deployer) getPreviousRequiredProviders(
	ctx context.Context,
	serviceID oid.ID,
) ([]types.ProviderRequirement, error) {
	prevRequiredProviders := make([]types.ProviderRequirement, 0)

	entity, err := d.modelClient.ServiceRevisions().Query().
		Where(servicerevision.ServiceID(serviceID)).
		Order(model.Desc(servicerevision.FieldCreateTime)).
		First(ctx)
	if err != nil && !model.IsNotFound(err) {
		return nil, err
	}

	if entity == nil {
		return prevRequiredProviders, nil
	}

	templateVersion, err := d.modelClient.TemplateVersions().Query().
		Where(
			templateversion.TemplateID(entity.TemplateID),
			templateversion.Version(entity.TemplateVersion),
		).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	if templateVersion.Schema != nil {
		prevRequiredProviders = append(prevRequiredProviders, templateVersion.Schema.RequiredProviders...)
	}

	return prevRequiredProviders, nil
}

// GetServiceDependencyOutputs gets the dependency outputs of the service.
func (d Deployer) getServiceDependencyOutputs(
	ctx context.Context,
	serviceID oid.ID,
	dependOutputs []string,
) (map[string]parser.OutputState, error) {
	service, err := d.modelClient.Services().Query().
		Where(service.ID(serviceID)).
		WithDependencies().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	dependantIDs := make([]oid.ID, 0, len(service.Edges.Dependencies))

	for _, d := range service.Edges.Dependencies {
		if d.Type != types.ServiceDependencyTypeImplicit {
			continue
		}

		dependantIDs = append(dependantIDs, d.DependentID)
	}

	dependencyRevisions, err := d.modelClient.ServiceRevisions().Query().
		Select(
			servicerevision.FieldID,
			servicerevision.FieldAttributes,
			servicerevision.FieldOutput,
			servicerevision.FieldServiceID,
			servicerevision.FieldProjectID,
		).
		Where(func(s *sql.Selector) {
			sq := s.Clone().
				AppendSelectExprAs(
					sql.RowNumber().
						PartitionBy(servicerevision.FieldServiceID).
						OrderBy(sql.Desc(servicerevision.FieldCreateTime)),
					"row_number",
				).
				Where(s.P()).
				From(s.Table()).
				As(servicerevision.Table)

				// Query the latest revision of the service.
			s.Where(sql.EQ(s.C("row_number"), 1)).
				From(sq)
		}).Where(servicerevision.ServiceIDIn(dependantIDs...)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	outputs := make(map[string]parser.OutputState, 0)
	dependSets := sets.NewString(dependOutputs...)

	for _, r := range dependencyRevisions {
		revisionOutput, err := parser.ParseStateOutputRawMap(r)
		if err != nil {
			return nil, err
		}

		for n, o := range revisionOutput {
			if dependSets.Has(n) {
				outputs[n] = o
			}
		}
	}

	return outputs, nil
}

func SyncServiceRevisionStatus(ctx context.Context, bm revisionbus.BusMessage) (err error) {
	var (
		mc       = bm.TransactionalModelClient
		revision = bm.Refer
	)

	// Report to service.
	service, err := mc.Services().Query().
		Where(service.ID(revision.ServiceID)).
		Select(
			service.FieldID,
			service.FieldStatus,
		).
		Only(ctx)
	if err != nil {
		return err
	}

	var serviceStatusUpdate *model.ServiceUpdateOne

	switch revision.Status {
	case status.ServiceRevisionStatusSucceeded:
		if status.ServiceStatusDeleted.IsUnknown(service) {
			err = mc.Services().DeleteOne(service).
				Exec(ctx)
		} else {
			status.ServiceStatusDeployed.True(service, "")
			status.ServiceStatusReady.Unknown(service, "")
			serviceStatusUpdate, err = dao.ServiceStatusUpdate(mc, service)
			if err != nil {
				return err
			}
			err = serviceStatusUpdate.Exec(ctx)
		}
	case status.ServiceRevisionStatusFailed:
		if status.ServiceStatusDeleted.IsUnknown(service) {
			status.ServiceStatusDeleted.False(service, "")
		} else {
			status.ServiceStatusDeployed.False(service, "")
		}
		service.Status.SummaryStatusMessage = revision.StatusMessage

		serviceStatusUpdate, err = dao.ServiceStatusUpdate(mc, service)
		if err != nil {
			return err
		}
		err = serviceStatusUpdate.Exec(ctx)
	}

	return err
}

// parseAttributeReplace parses attribute secrets ${secret.name} replaces it with ${var._secretPrefix+name},
// service reference ${service.name.output} replaces it with ${var._servicePrefix+name}
// and returns secret names and service names.
func parseAttributeReplace(
	attributes map[string]interface{},
	secretNames []string,
	serviceOutputs []string,
) ([]string, []string) {
	for key, value := range attributes {
		if value == nil {
			continue
		}

		switch reflect.TypeOf(value).Kind() {
		case reflect.Map:
			if _, ok := value.(map[string]interface{}); !ok {
				continue
			}

			secretNames, serviceOutputs = parseAttributeReplace(
				value.(map[string]interface{}),
				secretNames,
				serviceOutputs,
			)
		case reflect.String:
			str := value.(string)
			matches := _secretReg.FindAllStringSubmatch(str, -1)
			serviceMatches := _serviceReg.FindAllStringSubmatch(str, -1)

			var matched []string

			for _, match := range matches {
				if len(match) > 1 {
					matched = append(matched, match[1])
				}
			}

			var serviceMatched []string

			for _, match := range serviceMatches {
				if len(match) > 1 {
					serviceMatched = append(serviceMatched, match[1]+"_"+match[2])
				}
			}

			secretNames = append(secretNames, matched...)
			secretRepl := "${var." + _secretPrefix + "${1}}"
			str = _secretReg.ReplaceAllString(str, secretRepl)

			serviceOutputs = append(serviceOutputs, serviceMatched...)
			serviceRepl := "${var." + _servicePrefix + "${1}_${2}}"

			attributes[key] = _serviceReg.ReplaceAllString(str, serviceRepl)
		case reflect.Slice:
			if _, ok := value.([]interface{}); !ok {
				continue
			}

			for _, v := range value.([]interface{}) {
				if _, ok := v.(map[string]interface{}); !ok {
					continue
				}
				secretNames, serviceOutputs = parseAttributeReplace(
					v.(map[string]interface{}),
					secretNames,
					serviceOutputs,
				)
			}
		}
	}

	return secretNames, serviceOutputs
}

func getVarConfigOptions(secrets model.Secrets, serviceOutputs map[string]parser.OutputState) config.CreateOptions {
	varsConfigOpts := config.CreateOptions{
		Attributes: map[string]interface{}{},
	}

	for _, v := range secrets {
		varsConfigOpts.Attributes[_secretPrefix+v.Name] = v.Value
	}

	// Setup service outputs.
	for n, v := range serviceOutputs {
		varsConfigOpts.Attributes[_servicePrefix+n] = v.Value
	}

	return varsConfigOpts
}

func getModuleConfig(
	revision *model.ServiceRevision,
	modVer *model.TemplateVersion,
	ops CreateSecretsOptions,
) (*config.ModuleConfig, error) {
	var (
		props              = make(property.Properties, len(revision.Attributes))
		typesWith          = revision.Attributes.TypesWith(modVer.Schema.Variables)
		sensitiveVariables = sets.Set[string]{}
	)

	for k, v := range revision.Attributes {
		props[k] = property.Property{
			Type:  typesWith[k],
			Value: v,
		}

		// Add sensitive from app attributes.
		val, err := v.MarshalJSON()
		if err != nil {
			return nil, err
		}

		if _secretReg.Match(val) {
			sensitiveVariables.Insert(k)
		}
	}

	attrs, err := props.TypedValues()
	if err != nil {
		return nil, err
	}

	mc := &config.ModuleConfig{
		Name:       revision.Edges.Service.Name,
		Source:     modVer.Source,
		Schema:     modVer.Schema,
		Attributes: attrs,
	}

	if modVer.Schema == nil {
		return mc, nil
	}

	for _, v := range modVer.Schema.Variables {
		// Add sensitive from schema variable.
		if v.Sensitive {
			sensitiveVariables.Insert(v.Name)
		}

		// Add seal metadata.
		var attrValue string

		switch v.Name {
		case SealMetadataProjectName:
			attrValue = ops.ProjectName
		case SealMetadataEnvironmentName:
			attrValue = ops.EnvironmentName
		case SealMetadataServiceName:
			attrValue = ops.ServiceName
		}

		if attrValue != "" {
			mc.Attributes[v.Name] = attrValue
		}
	}

	sensitiveVariableRegex, err := matchAnyRegex(sensitiveVariables.UnsortedList())
	if err != nil {
		return nil, err
	}

	mc.Outputs = make([]config.Output, len(modVer.Schema.Outputs))
	for i, v := range modVer.Schema.Outputs {
		mc.Outputs[i].Sensitive = v.Sensitive
		mc.Outputs[i].Name = v.Name
		mc.Outputs[i].ServiceName = revision.Edges.Service.Name

		if v.Sensitive {
			continue
		}

		// Update sensitive while output is from sensitive data, like secret.
		if sensitiveVariables.Len() != 0 && sensitiveVariableRegex.Match(v.Value) {
			mc.Outputs[i].Sensitive = true
		}
	}

	return mc, nil
}

func matchAnyRegex(list []string) (*regexp.Regexp, error) {
	var sb strings.Builder

	sb.WriteString("(")

	for i, v := range list {
		sb.WriteString(fmt.Sprintf(`var\.%s`, v))

		if i < len(list)-1 {
			sb.WriteString("|")
		}
	}

	sb.WriteString(")")

	return regexp.Compile(sb.String())
}