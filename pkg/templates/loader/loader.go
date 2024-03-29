package loader

import (
	"fmt"

	"github.com/seal-io/walrus/pkg/templates/loaders/terraform"
	"github.com/seal-io/walrus/pkg/types"
)

// SchemaLoader define the interface for loading schema from template.
type SchemaLoader interface {
	Load(rootDir, templateName string) (*types.SchemaGroup, error)
}

// LoadSchema loads schema from template.
func LoadSchema(rootDir, templateName, templateFormat string) (s *types.SchemaGroup, err error) {
	// Terraform.
	switch templateFormat {
	default:
		return nil, fmt.Errorf("unsupport template format %s", templateFormat)
	case types.TemplateFormatTerraform:
		tf := terraform.New()
		return tf.Load(rootDir, templateName)
	}
}
