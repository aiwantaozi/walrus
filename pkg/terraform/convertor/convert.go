package convertor

import (
	"fmt"

	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/terraform/block"
	"github.com/seal-io/walrus/utils/strs"
)

const (
	TypeKubernetes = "kubernetes"
	TypeHelm       = "helm"
	TypeKubectl    = "kubectl"
	TypeAliCloud   = "alicloud"
	TypeAWS        = "aws"
	TypeDocker     = "docker"
	// Add more convertor type.
)

// supportedProviders is the supported provider types,
// it is used to validate the required providers.
var supportedProviders = []string{
	TypeKubernetes,
	TypeHelm,
	TypeKubectl,
	TypeAliCloud,
	TypeAWS,
	// Add more convertor type.
}

type ConvertOptions struct {
	SecretMountPath string
	ConnSeparator   string
	Providers       []string
}

// LoadConvertor loads the convertor by the provider type.
func LoadConvertor(provider string) Convertor {
	switch provider {
	case TypeKubernetes:
		return K8sConvertor(provider)
	case TypeHelm:
		return HelmConvertor(provider)
	case TypeKubectl:
		return KubectlConvertor(provider)
	case TypeAliCloud:
		return AlibabaConvertor(provider)
	case TypeAWS:
		return AWSConvertor(provider)
	case TypeDocker:
		return DockerConvertor(provider)
	default:
		return DefaultConvertor(provider)
	}
}

// ToProvidersBlocks converts the connectors to provider blocks with required providers.
func ToProvidersBlocks(
	providers []string,
	connectors model.Connectors,
	opts ConvertOptions,
) (blocks block.Blocks, err error) {
	for _, p := range providers {
		var convertBlocks block.Blocks

		convertBlocks, err = ToProviderBlocks(p, connectors, opts)
		if err != nil {
			return nil, err
		}

		if convertBlocks == nil {
			continue
		}

		blocks = append(blocks, convertBlocks...)
	}

	// Validate blocks with providers.
	ok, err := validateRequiredProviders(providers, blocks)
	if err != nil {
		return nil, err
	}

	currentProviders, err := blocks.GetProviderNames()
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, fmt.Errorf(
			"missing required providers, expected: %q, but got %q",
			strs.Join(", ", providers...),
			strs.Join(", ", currentProviders...),
		)
	}

	return blocks, nil
}

// ToProviderBlocks converts the connectors to blocks with provider name.
func ToProviderBlocks(provider string, connectors model.Connectors, opts ConvertOptions) (block.Blocks, error) {
	var toBlockOpts Options

	switch provider {
	case TypeKubernetes, TypeHelm, TypeKubectl:
		toBlockOpts = K8sConvertorOptions{
			ConfigPath:    opts.SecretMountPath,
			ConnSeparator: opts.ConnSeparator,
		}
	default:
		toBlockOpts = opts
	}

	return LoadConvertor(provider).ToBlocks(connectors, toBlockOpts)
}

// validateRequiredProviders validate the required providers in the generated blocks.
func validateRequiredProviders(providers []string, blocks block.Blocks) (bool, error) {
	// BlockProviders is the providers of the generated blocks.
	blockProviders, err := blocks.GetProviderNames()
	if err != nil {
		return false, err
	}
	actual := sets.NewString(blockProviders...)

	// Supported is the providers which convertors can generate with the given connectors.
	// Custom providers are the providers which are generated by the primary provider.
	supported := sets.NewString(supportedProviders...)

	// Get the intersection of the required and supported providers to validate.
	required := sets.NewString(providers...)
	expected := required.Intersection(supported)

	return actual.IsSuperset(expected), nil
}
