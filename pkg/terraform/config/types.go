package config

import (
	"github.com/hashicorp/terraform-config-inspect/tfconfig"

	"github.com/seal-io/seal/pkg/dao/model"
	"github.com/seal-io/seal/pkg/dao/types"
	"github.com/seal-io/seal/pkg/terraform/parser"
)

const (
	FileMain = "main.tf"
	FileVars = "terraform.tfvars"
)

// ModuleConfig is a struct with model.Template and its variables.
type ModuleConfig struct {
	// Name is the module name.
	Name string
	// Source is the module source.
	Source string
	// Schema is the variable schema.
	Schema *types.TemplateSchema
	// Attributes is the attributes of the module.
	Attributes map[string]interface{}
	// Outputs is the module outputs.
	Outputs []Output
}

// CreateOptions represents the CreateOptions of the Config.
type CreateOptions struct {
	Attributes       map[string]interface{}
	TerraformOptions *TerraformOptions
	ProviderOptions  *ProviderOptions
	ModuleOptions    *ModuleOptions
	VariableOptions  *VariableOptions
	OutputOptions    OutputOptions
}

type (
	// TerraformOptions is the options to create terraform block.
	TerraformOptions struct {
		// Token is the backend token to authenticate with the Seal Server of the terraform config.
		Token string
		// Address is the backend address of the terraform config.
		Address string
		// SkipTLSVerify is the backend cert verification of the terraform config.
		SkipTLSVerify bool

		// ProviderRequirements is the required providers of the terraform config.
		ProviderRequirements map[string]*tfconfig.ProviderRequirement
	}

	// ProviderOptions is the options to create provider blocks.
	ProviderOptions struct {
		// SecretMountPath is the mount path of the secret in the terraform config.
		SecretMonthPath string
		// ConnectorSeparator is the separator of the terraform provider alias name and id.
		ConnectorSeparator string
		// RequiredProviderNames is the required providers of the terraform config.
		// E.g. ["kubernetes", "helm"].
		RequiredProviderNames []string
		Connectors            model.Connectors
	}

	// ModuleOptions is the options to create module blocks.
	ModuleOptions struct {
		// ModuleConfigs is the module configs of the deployment.
		ModuleConfigs []*ModuleConfig
	}

	// VariableOptions is the options to create variables blocks.
	VariableOptions struct {
		// SecretPrefix is the prefix of the secret name.
		SecretPrefix string
		// SecretServicePrefix is the prefix of the secret service name.
		ServicePrefix string
		// SecretNames is the list of the secret name, type is always string.
		SecretNames []string
		// DependencyOutputs is the map of the variable name and value.
		DependencyOutputs map[string]parser.OutputState
	}

	// OutputOptions is the options to create outputs blocks.
	OutputOptions []Output
	// Output indicate the output name and module.
	Output struct {
		ServiceName string
		Name        string
		Sensitive   bool
	}
)