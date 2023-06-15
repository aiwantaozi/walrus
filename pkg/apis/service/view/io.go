package view

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"k8s.io/utils/pointer"

	"github.com/seal-io/seal/pkg/apis/runtime"
	"github.com/seal-io/seal/pkg/dao/model"
	"github.com/seal-io/seal/pkg/dao/model/environment"
	"github.com/seal-io/seal/pkg/dao/model/environmentconnectorrelationship"
	"github.com/seal-io/seal/pkg/dao/model/predicate"
	"github.com/seal-io/seal/pkg/dao/model/service"
	"github.com/seal-io/seal/pkg/dao/model/servicedependency"
	"github.com/seal-io/seal/pkg/dao/model/serviceresource"
	"github.com/seal-io/seal/pkg/dao/model/servicerevision"
	"github.com/seal-io/seal/pkg/dao/model/templateversion"
	"github.com/seal-io/seal/pkg/dao/types"
	"github.com/seal-io/seal/pkg/dao/types/oid"
	"github.com/seal-io/seal/pkg/dao/types/status"
	"github.com/seal-io/seal/pkg/topic/datamessage"
	"github.com/seal-io/seal/utils/validation"
)

// Basic APIs.

type CreateRequest struct {
	*model.ServiceCreateInput `json:",inline"`

	ProjectID  oid.ID   `query:"projectID"`
	RemarkTags []string `json:"remarkTags,omitempty"`
}

func (r *CreateRequest) ValidateWith(ctx context.Context, input any) error {
	if !r.ProjectID.Valid(0) {
		return errors.New("invalid project id: blank")
	}

	if !r.Environment.ID.Valid(0) {
		return errors.New("invalid environment id: blank")
	}

	if err := validation.IsDNSSubdomainName(r.Name); err != nil {
		return fmt.Errorf("invalid name: %w", err)
	}

	// Verify module version.
	modelClient := input.(model.ClientSet)

	templateVersion, err := modelClient.TemplateVersions().Query().
		Where(templateversion.TemplateID(r.Template.ID), templateversion.Version(r.Template.Version)).
		Only(ctx)
	if err != nil {
		return runtime.Errorw(err, "failed to get template version")
	}

	// Verify environment if it has no connectors.
	_, err = modelClient.Environments().Query().
		Where(environment.ID(r.Environment.ID)).
		OnlyID(ctx)
	if err != nil {
		return runtime.Errorw(err, "failed to get environment")
	}

	count, _ := modelClient.EnvironmentConnectorRelationships().Query().
		Where(environmentconnectorrelationship.EnvironmentID(r.Environment.ID)).
		Count(ctx)
	if count == 0 {
		return runtime.Error(http.StatusNotFound, "invalid environment: no connectors")
	}

	// Verify variables with variables schema that defined on the template version.
	err = r.Attributes.ValidateWith(templateVersion.Schema.Variables)
	if err != nil {
		return fmt.Errorf("invalid variables: %w", err)
	}

	return nil
}

type CreateResponse = *model.ServiceOutput

type DeleteRequest struct {
	*model.ServiceQueryInput `uri:",inline"`

	ProjectID oid.ID `query:"projectID"`
	Force     *bool  `query:"force,default=true"`
}

func (r *DeleteRequest) ValidateWith(ctx context.Context, input any) error {
	if !r.ProjectID.Valid(0) {
		return errors.New("invalid project id: blank")
	}

	if !r.ID.Valid(0) {
		return errors.New("invalid id: blank")
	}

	modelClient := input.(model.ClientSet)

	ids, err := modelClient.ServiceDependencies().Query().
		Where(servicedependency.DependentID(r.ID)).
		IDs(ctx)
	if err != nil {
		return runtime.Errorw(err, "failed to get service dependencies")
	}

	if len(ids) > 0 {
		return runtime.Error(http.StatusConflict, "service has dependencies")
	}

	if r.Force == nil {
		// By default, clean deployed native resources too.
		r.Force = pointer.Bool(true)
	}

	if *r.Force {
		err := validateRevisionStatus(ctx, modelClient, r.ID, "delete")
		if err != nil {
			return err
		}
	}

	return nil
}

type GetRequest struct {
	*model.ServiceQueryInput `uri:",inline"`

	ProjectID oid.ID `query:"projectID"`
}

func (r *GetRequest) Validate() error {
	if !r.ProjectID.Valid(0) {
		return errors.New("invalid project id: blank")
	}

	if !r.ID.Valid(0) {
		return errors.New("invalid id: blank")
	}

	return nil
}

type GetResponse = *model.ServiceOutput

type StreamRequest struct {
	ID oid.ID `uri:"id"`

	ProjectID oid.ID `query:"projectID"`
}

func (r *StreamRequest) ValidateWith(ctx context.Context, input any) error {
	if !r.ProjectID.Valid(0) {
		return errors.New("invalid project id: blank")
	}

	if !r.ID.Valid(0) {
		return errors.New("invalid id: blank")
	}

	modelClient := input.(model.ClientSet)

	exist, err := modelClient.Services().Query().
		Where(service.ID(r.ID)).
		Exist(ctx)
	if err != nil || !exist {
		return runtime.Errorw(err, "invalid id: not found")
	}

	return nil
}

type StreamResponse struct {
	Type       datamessage.EventType  `json:"type"`
	IDs        []oid.ID               `json:"ids,omitempty"`
	Collection []*model.ServiceOutput `json:"collection,omitempty"`
}

// Batch APIs.

type CollectionGetRequest struct {
	runtime.RequestCollection[predicate.Service, service.OrderOption] `query:",inline"`

	ProjectID     oid.ID `query:"projectID"`
	EnvironmentID oid.ID `query:"environmentID"`
}

func (r *CollectionGetRequest) ValidateWith(ctx context.Context, input any) error {
	if !r.ProjectID.Valid(0) {
		return errors.New("invalid project id: blank")
	}

	if !r.EnvironmentID.Valid(0) {
		return errors.New("invalid environment id: blank")
	}

	modelClient := input.(model.ClientSet)

	_, err := modelClient.Environments().Query().
		Where(environment.ID(r.EnvironmentID)).
		OnlyID(ctx)
	if err != nil {
		return runtime.Errorw(err, "failed to get environment")
	}

	return nil
}

type CollectionGetResponse = []*model.ServiceOutput

type CollectionStreamRequest struct {
	runtime.RequestExtracting `query:",inline"`

	ProjectID     oid.ID `query:"projectID"`
	EnvironmentID oid.ID `query:"environmentID,omitempty"`
}

func (r *CollectionStreamRequest) ValidateWith(ctx context.Context, input any) error {
	if !r.ProjectID.Valid(0) {
		return errors.New("invalid project id: blank")
	}

	if r.EnvironmentID.Valid(0) {
		modelClient := input.(model.ClientSet)

		if !r.EnvironmentID.Valid(0) {
			return errors.New("invalid environment id: blank")
		}

		_, err := modelClient.Environments().Query().
			Where(environment.ID(r.EnvironmentID)).
			OnlyID(ctx)
		if err != nil {
			return runtime.Errorw(err, "failed to get environment")
		}
	}

	return nil
}

// Extensional APIs.

type RouteUpgradeRequest struct {
	_ struct{} `route:"PUT=/upgrade"`

	*model.ServiceUpdateInput `uri:",inline" json:",inline"`

	ProjectID  oid.ID   `query:"projectID"`
	RemarkTags []string `json:"remarkTags,omitempty"`
}

func (r *RouteUpgradeRequest) ValidateWith(ctx context.Context, input any) error {
	if !r.ProjectID.Valid(0) {
		return errors.New("invalid project id: blank")
	}

	if !r.ID.Valid(0) {
		return errors.New("invalid id: blank")
	}

	modelClient := input.(model.ClientSet)

	service, err := modelClient.Services().Query().
		Select(
			service.FieldTemplate,
		).
		Where(service.ID(r.ID)).
		Only(ctx)
	if err != nil {
		return runtime.Errorw(err, "failed to get service")
	}

	if r.Template.ID != service.Template.ID {
		return errors.New("invalid template id: immutable")
	}

	tv, err := modelClient.TemplateVersions().Query().
		Select(templateversion.FieldSchema).
		Where(
			templateversion.TemplateID(r.Template.ID),
			templateversion.Version(r.Template.Version),
		).
		Only(ctx)
	if err != nil {
		return runtime.Errorw(err, "failed to get template version")
	}

	// Verify attributes with variables schema of the template version.
	err = r.Attributes.ValidateWith(tv.Schema.Variables)
	if err != nil {
		return fmt.Errorf("invalid variables: %w", err)
	}

	err = validateRevisionStatus(ctx, modelClient, r.ID, "upgrade")
	if err != nil {
		return err
	}

	return nil
}

func IsEndpointOutput(outputName string) bool {
	return strings.HasPrefix(outputName, "endpoint")
}

type AccessEndpointRequest struct {
	_ struct{} `route:"GET=/access-endpoints"`

	*model.ServiceQueryInput `uri:",inline"`

	ProjectID oid.ID `query:"projectID"`
}

func (r *AccessEndpointRequest) ValidateWith(ctx context.Context, input any) error {
	if !r.ProjectID.Valid(0) {
		return errors.New("invalid project id: blank")
	}

	if !r.ID.Valid(0) {
		return errors.New("invalid id: blank")
	}

	modelClient := input.(model.ClientSet)

	_, err := modelClient.Services().Query().
		Where(service.ID(r.ID)).
		OnlyID(ctx)
	if err != nil {
		return runtime.Errorw(err, "failed to get service")
	}

	return nil
}

type AccessEndpointResponse = []Endpoint

type Endpoint struct {
	// Name is identifier for the endpoint.
	Name string `json:"name,omitempty"`
	// Endpoint is access endpoint.
	Endpoints []string `json:"endpoints,omitempty"`
}

type OutputRequest struct {
	_ struct{} `route:"GET=/outputs"`

	*model.ServiceQueryInput `uri:",inline"`

	ProjectID oid.ID `query:"projectID"`
}

func (r *OutputRequest) ValidateWith(ctx context.Context, input any) error {
	if !r.ProjectID.Valid(0) {
		return errors.New("invalid project id: blank")
	}

	if !r.ID.Valid(0) {
		return errors.New("invalid id: blank")
	}

	modelClient := input.(model.ClientSet)

	_, err := modelClient.Services().Query().
		Where(service.ID(r.ID)).
		OnlyID(ctx)
	if err != nil {
		return runtime.Errorw(err, "failed to get service")
	}

	return nil
}

type OutputResponse = []types.OutputValue

type CreateCloneRequest struct {
	_ struct{} `route:"POST=/clone"`

	ID            oid.ID   `uri:"id"`
	EnvironmentID oid.ID   `json:"environmentID"`
	Name          string   `json:"name"`
	RemarkTags    []string `json:"remarkTags,omitempty"`
}

func (r *CreateCloneRequest) ValidateWith(ctx context.Context, input any) error {
	if !r.ID.Valid(0) {
		return errors.New("invalid id: blank")
	}

	if r.Name == "" {
		return errors.New("invalid name: blank")
	}

	if r.EnvironmentID != "" {
		if !r.EnvironmentID.IsNaive() {
			return fmt.Errorf("invalid environment id: %s", r.EnvironmentID)
		}
		modelClient := input.(model.ClientSet)

		_, err := modelClient.Environments().Query().
			Where(environment.ID(r.EnvironmentID)).
			OnlyID(ctx)
		if err != nil {
			return runtime.Errorw(err, "failed to get environment")
		}
	}

	return nil
}

func validateRevisionStatus(
	ctx context.Context,
	modelClient model.ClientSet,
	id oid.ID,
	action string,
) error {
	revision, err := modelClient.ServiceRevisions().Query().
		Where(servicerevision.ServiceID(id)).
		Order(model.Desc(servicerevision.FieldCreateTime)).
		First(ctx)
	if err != nil && !model.IsNotFound(err) {
		return runtime.Errorw(err, "failed to get service revision")
	}

	if revision != nil {
		switch revision.Status {
		case status.ServiceRevisionStatusSucceeded:
		case status.ServiceRevisionStatusRunning:
			return runtime.Error(http.StatusBadRequest,
				"deployment is running, please wait for it to finish before deleting the service")
		case status.ServiceRevisionStatusFailed:
			if action != "delete" {
				return nil
			}

			resourceExist, err := modelClient.ServiceResources().Query().
				Where(serviceresource.ServiceID(id)).
				Exist(ctx)
			if err != nil {
				return err
			}

			if resourceExist {
				return runtime.Error(
					http.StatusBadRequest,
					"latest deployment is not succeeded,"+
						" please fix the service configuration or rollback before deleting it",
				)
			}
		default:
			return runtime.Error(http.StatusBadRequest, "invalid deployment status")
		}
	}

	return nil
}

type StreamAccessEndpointResponse struct {
	Type       datamessage.EventType  `json:"type"`
	Collection AccessEndpointResponse `json:"collection,omitempty"`
}

type StreamOutputResponse struct {
	Type       datamessage.EventType `json:"type"`
	Collection OutputResponse        `json:"collection,omitempty"`
}