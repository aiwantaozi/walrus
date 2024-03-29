package types

import (
	"context"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/hashicorp/terraform-config-inspect/tfconfig"
	"helm.sh/helm/v3/pkg/chart"

	"github.com/seal-io/utils/json"
	//"github.com/seal-io/walrus/utils/log"

	"github.com/seal-io/walrus/pkg/templates/openapi"
)

const (
	VariableSchemaKey = "variables"
	OutputSchemaKey   = "outputs"
)

const (
	TagUserEdited = "user-edited"
)

type SchemaGroup struct {
	Schema   *Schema   `json:"schema"`
	UISchema *UISchema `json:"uiSchema"`
}

// OpenAPISchema specifies the openAPI schema with variables and outputs.
type OpenAPISchema struct {
	OpenAPISchema *openapi3.T `json:"openAPISchema"`
}

// Validate reports if the schema is valid.
func (s *OpenAPISchema) Validate() error {
	if s.OpenAPISchema == nil {
		return nil
	}

	// workaround: inject paths and version since kin-openapi/openapi3 need it.
	s.OpenAPISchema.Paths = &openapi3.Paths{}
	if s.OpenAPISchema.Info != nil && s.OpenAPISchema.Info.Version == "" {
		s.OpenAPISchema.Info.Version = "v0.0.0"
	}

	if err := s.OpenAPISchema.Validate(
		context.Background(),
		openapi3.DisableSchemaDefaultsValidation(),
	); err != nil {
		return err
	}

	return nil
}

func (s *OpenAPISchema) IsEmpty() bool {
	return s.OpenAPISchema == nil || s.OpenAPISchema.Components == nil || len(s.OpenAPISchema.Components.Schemas) == 0
}

// Expose returns the UI schema of the schema.
func (s *OpenAPISchema) Expose(skipProps ...string) UISchema {
	vs := s.VariableSchema()
	if vs == nil {
		return UISchema{}
	}

	// In order to prevent the remove ext affect the original schema, serialize and deserialize to copy the schema.
	b, err := json.Marshal(vs)
	if err != nil {
		// TODO(michelia): log
		//log.Warnf("error marshal variable schema while expost: %v", err)
		return UISchema{}
	}

	var cps openapi3.Schema

	err = json.Unmarshal(b, &cps)
	if err != nil {
		// TODO(michelia): log
		//log.Warnf("error unmarshal variable schema while expost: %v", err)
		return UISchema{}
	}

	for _, v := range skipProps {
		delete(cps.Properties, v)
	}

	return UISchema{
		OpenAPISchema: &openapi3.T{
			OpenAPI: s.OpenAPISchema.OpenAPI,
			Info:    s.OpenAPISchema.Info,
			Components: &openapi3.Components{
				Schemas: map[string]*openapi3.SchemaRef{
					VariableSchemaKey: {
						Value: openapi.RemoveExtOriginal(&cps),
					},
				},
			},
		},
	}
}

// VariableSchema returns the variables' schema.
func (s *OpenAPISchema) VariableSchema() *openapi3.Schema {
	if s.OpenAPISchema == nil ||
		s.OpenAPISchema.Components == nil ||
		s.OpenAPISchema.Components.Schemas == nil ||
		s.OpenAPISchema.Components.Schemas[VariableSchemaKey] == nil ||
		s.OpenAPISchema.Components.Schemas[VariableSchemaKey].Value == nil {
		return nil
	}

	return s.OpenAPISchema.Components.Schemas[VariableSchemaKey].Value
}

func (s *OpenAPISchema) SetVariableSchema(v *openapi3.Schema) {
	s.ensureInit()
	s.OpenAPISchema.Components.Schemas[VariableSchemaKey].Value = v
}

func (s *OpenAPISchema) RemoveVariableContext() {
	if s.IsEmpty() {
		return
	}

	variableSchema := openapi.RemoveVariableContext(s.VariableSchema())
	s.SetVariableSchema(variableSchema)
}

func (s *OpenAPISchema) SetOutputSchema(v *openapi3.Schema) {
	s.ensureInit()
	s.OpenAPISchema.Components.Schemas[OutputSchemaKey].Value = v
}

func (s *OpenAPISchema) ensureInit() {
	if s.OpenAPISchema == nil {
		s.OpenAPISchema = &openapi3.T{}
	}

	if s.OpenAPISchema.Components == nil {
		s.OpenAPISchema.Components = &openapi3.Components{}
	}

	if s.OpenAPISchema.Components.Schemas == nil {
		s.OpenAPISchema.Components.Schemas = openapi3.Schemas{}
	}
}

// OutputSchema returns the outputs' schema.
func (s *OpenAPISchema) OutputSchema() *openapi3.Schema {
	if s.OpenAPISchema == nil ||
		s.OpenAPISchema.Components == nil ||
		s.OpenAPISchema.Components.Schemas == nil ||
		s.OpenAPISchema.Components.Schemas[OutputSchemaKey] == nil ||
		s.OpenAPISchema.Components.Schemas[OutputSchemaKey].Value == nil {
		return nil
	}

	return s.OpenAPISchema.Components.Schemas[OutputSchemaKey].Value
}

// Intersect sets variables & outputs schema of s to intersection of s and s2.
func (s *OpenAPISchema) Intersect(s2 *OpenAPISchema) {
	if s2.OpenAPISchema == nil {
		return
	}

	variableSchema := openapi.IntersectSchema(s.VariableSchema(), s2.VariableSchema())
	s.SetVariableSchema(variableSchema)
	outputSchema := openapi.IntersectSchema(s.OutputSchema(), s2.OutputSchema())
	s.SetOutputSchema(outputSchema)
}

// UISchema include the UI schema that users can customize.
type UISchema OpenAPISchema

// IsEmpty reports if the schema is empty.
func (s UISchema) IsEmpty() bool {
	return s.OpenAPISchema == nil ||
		s.OpenAPISchema.Components == nil ||
		len(s.OpenAPISchema.Components.Schemas) == 0 ||
		s.OpenAPISchema.Components.Schemas[VariableSchemaKey] == nil ||
		s.OpenAPISchema.Components.Schemas[VariableSchemaKey].Value == nil
}

// IsUserEdited reports if the ui schema is edited by user.
func (s *UISchema) IsUserEdited() bool {
	return s.OpenAPISchema != nil && s.OpenAPISchema.Tags.Get(TagUserEdited) != nil
}

// VariableSchema returns the variables' schema.
func (s *UISchema) VariableSchema() *openapi3.Schema {
	if s.IsEmpty() {
		return nil
	}

	return s.OpenAPISchema.Components.Schemas[VariableSchemaKey].Value
}

// SetVariableSchema sets the variables' schema.
func (s *UISchema) SetVariableSchema(v *openapi3.Schema) {
	if s.OpenAPISchema == nil {
		return
	}

	s.OpenAPISchema.Components.Schemas[VariableSchemaKey].Value = v
}

// SetUserEdited sets the user edited tag to the ui schema.
func (s *UISchema) SetUserEdited() {
	if s.OpenAPISchema != nil {
		tag := s.OpenAPISchema.Tags.Get(TagUserEdited)
		if tag == nil {
			s.OpenAPISchema.Tags = append(
				s.OpenAPISchema.Tags,
				&openapi3.Tag{Name: TagUserEdited})
		}
	}
}

// UnsetUserEdited unsets the user edited tag to the ui schema.
func (s *UISchema) UnsetUserEdited() {
	if s.OpenAPISchema != nil {
		tag := s.OpenAPISchema.Tags.Get(TagUserEdited)
		if tag == nil {
			return
		}

		newTags := make(openapi3.Tags, len(s.OpenAPISchema.Tags)-1)
		for _, v := range s.OpenAPISchema.Tags {
			if v.Name != TagUserEdited {
				newTags = append(newTags, v)
			}
		}

		s.OpenAPISchema.Tags = newTags
	}
}

// Validate reports if the ui schema is valid.
func (s *UISchema) Validate() error {
	w := OpenAPISchema{
		OpenAPISchema: s.OpenAPISchema,
	}

	return w.Validate()
}

type TemplateData struct {
	// Readme specifies the readme of this template.
	Readme string `json:"readme"`

	Helm      HelmMetadata      `json:"helm,omitempty"`
	Terraform TerraformMetadata `json:"terraform,omitempty"`
}

// Schema include the internal template variables schema and template data.
type Schema struct {
	OpenAPISchema `json:",inline"`

	// SchemaDefaultValue specifies the default value of the schema.
	SchemaDefaultValue []byte `json:"schemaDefaultValue"`

	// Metadata specifies the metadata of this template.
	Data TemplateData `json:"data"`
}

// Validate reports if the schema is valid.
func (s *Schema) Validate() error {
	return s.OpenAPISchema.Validate()
}

// TerraformMetadata include the terraform metadata of this template version.
type TerraformMetadata struct {
	// RequiredProviders specifies the required providers of this template.
	RequiredProviders []ProviderRequirement `json:"requiredProviders"`
}

// ProviderRequirement include the required provider.
type ProviderRequirement struct {
	*tfconfig.ProviderRequirement

	Name string `json:"name"`
}

type HelmMetadata struct {
	Chart chart.Metadata `json:"chart"`
}
