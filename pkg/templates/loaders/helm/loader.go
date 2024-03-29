package helm

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/hashicorp/terraform-config-inspect/tfconfig"
	"gopkg.in/yaml.v3"
	"helm.sh/helm/v3/pkg/chart"
	chartloader "helm.sh/helm/v3/pkg/chart/loader"

	"github.com/seal-io/walrus/pkg/templates/translator"
	"github.com/seal-io/walrus/pkg/templates/translators/terraform"

	"github.com/seal-io/utils/files"
	"github.com/seal-io/walrus/pkg/templates/openapi"
)

const (
	defaultGroup = "Basic"
)

var (
	readmeFileNames = []string{"readme.md", "readme.txt", "readme"}
	valuesFileNames = []string{"values.yaml", "values.yml"}
)

// Loader for load terraform template schema and data.
type Loader struct {
	translator translator.Translator
}

// New creates a new terraform loader.
func New() *Loader {
	return &Loader{
		translator: terraform.New(),
	}
}

// Load loads the internal template version schema and data.
func (l *Loader) Load(
	rootDir, templateName string,
) (*mytypes.SchemaGroup, error) {
	chart, err := l.loadChart(rootDir)
	if err != nil {
		return nil, err
	}

	ts, fs, err := l.loadSchema(rootDir, chart, templateName)
	if err != nil {
		return nil, err
	}

	return &mytypes.SchemaGroup{
		Schema:   ts,
		UISchema: fs,
	}, nil
}

// loadChart load the helm chart.
func (l *Loader) loadChart(rootDir string) (*chart.Chart, error) {
	return chartloader.LoadDir(rootDir)
}

// loadSchema loads the internal template version schema.
func (l *Loader) loadSchema(
	rootDir string,
	ct *chart.Chart,
	template string,
) (*mytypes.Schema, *mytypes.UISchema, error) {
	// OpenAPISchema.
	ts, err := l.getSchema(ct, template)
	if err != nil {
		return nil, nil, err
	}

	openAPISchema := mytypes.OpenAPISchema{
		OpenAPISchema: ts,
	}

	tsDefault, err := openapi.GenSchemaDefaultPatch(context.Background(), openAPISchema.VariableSchema())
	if err != nil {
		return nil, nil, err
	}

	d, err := l.loadData(ct)
	if err != nil {
		return nil, nil, err
	}

	wts := &mytypes.Schema{
		OpenAPISchema:      openAPISchema,
		SchemaDefaultValue: tsDefault,
		Data:               d,
	}

	// UI OpenAPISchema.
	fs, err := l.getSchemaFromFile(rootDir, ts)
	if err != nil {
		return nil, nil, err
	}

	wfs := mytypes.Schema{
		OpenAPISchema: mytypes.OpenAPISchema{
			OpenAPISchema: fs,
		},
	}

	uiSchema := wts.Expose(openapi.WalrusContextVariableName)
	if fs != nil && !wfs.IsEmpty() {
		uiSchema = wfs.Expose()
	}

	return wts, &uiSchema, nil
}

func (l *Loader) getSchema(chart *chart.Chart, template string) (*openapi3.T, error) {
	varsSchema, err := l.getVariableSchema(chart)
	if err != nil {
		return nil, err
	}

	// OpenAPI OpenAPISchema.
	t := &openapi3.T{
		OpenAPI: openapi.OpenAPIVersion,
		Info: &openapi3.Info{
			Title: fmt.Sprintf("OpenAPI schema for template %s", template),
		},
		Components: &openapi3.Components{
			Schemas: openapi3.Schemas{},
		},
	}

	if varsSchema != nil {
		t.Components.Schemas["variables"] = varsSchema.NewRef()
	}

	return t, nil
}

// getSchemaFromFile get openapi schema from schema.yaml.
func (l *Loader) getSchemaFromFile(rootDir string, originalSchema *openapi3.T) (*openapi3.T, error) {
	schemaFile := filepath.Join(rootDir, "schema.yaml")
	if !files.Exists(schemaFile) {
		if schemaFile = filepath.Join(rootDir, "schema.yml"); !files.Exists(schemaFile) {
			return nil, nil
		}
	}

	content, err := os.ReadFile(schemaFile)
	if err != nil {
		return nil, fmt.Errorf("error reading schema file %s: %w", schemaFile, err)
	}

	if len(content) == 0 {
		return nil, nil
	}

	// Openapi loader will cache the data with file path as key if we use LoadFromFile,
	// since the repo with different tag the schema.yaml file is the same, so we use LoadFromData to skip the cache.
	it, err := openapi3.NewLoader().LoadFromData(content)
	if err != nil {
		return nil, fmt.Errorf("error loading schema file %s: %w", schemaFile, err)
	}

	if it.Components == nil ||
		it.Components.Schemas == nil ||
		it.Components.Schemas["variables"] == nil ||
		it.Components.Schemas["variables"].Value == nil {
		return nil, nil
	}

	// Inject.
	var (
		varsSchema         = it.Components.Schemas["variables"].Value
		originalVarsSchema *openapi3.Schema
	)

	if originalSchema.Components != nil && originalSchema.Components.Schemas["variables"] != nil {
		originalVarsSchema = originalSchema.Components.Schemas["variables"].Value
	}

	varsSchema = l.applyMissingConfig(originalVarsSchema, varsSchema)
	l.injectExts(varsSchema)
	it.Components.Schemas["variables"].Value = varsSchema

	return it, nil
}

// applyMissingConfig apply the missing config to schema generate from schema.yaml.
func (l *Loader) applyMissingConfig(generated, customized *openapi3.Schema) *openapi3.Schema {
	if customized == nil {
		return nil
	}

	if generated == nil {
		return customized
	}

	s := *customized
	if len(s.Extensions) == 0 && len(generated.Extensions) != 0 {
		s.Extensions = generated.Extensions
	}

	for n, v := range s.Properties {
		in := generated.Properties[n]
		if in == nil || in.Value == nil {
			continue
		}

		// Title.
		if v.Value.Title == "" {
			s.Properties[n].Value.Title = generated.Properties[n].Value.Title
		}

		// Extensions.
		var (
			genExt = openapi.NewExtFromMap(in.Value.Extensions)
			ext    = openapi.NewExtFromMap(v.Value.Extensions)
		)

		ext.WithOriginal(in.Value.Extensions[openapi.ExtOriginalKey])

		if ext.ExtUI.Order == 0 {
			ext.WithUIOrder(genExt.Order)
		}

		if ext.ExtUI.ColSpan == 0 {
			ext.WithUIColSpan(genExt.ColSpan)
		}

		s.Properties[n].Value.Extensions = ext.Export()
	}

	return &s
}

// getVariableSchema generate variable schemas from helm values file.
func (l *Loader) getVariableSchema(ct *chart.Chart) (*openapi3.Schema, error) {
	if len(ct.Values) == 0 {
		return nil, nil
	}

	var (
		varSchemas = openapi3.NewObjectSchema()
		required   []string
		keys       = make([]string, len(ct.Values))
	)

	file := findFile(ct, valuesFileNames)
	if file == nil || len(file.Data) == 0 {
		return nil, fmt.Errorf("values file not found")
	}

	content, err := FormatLineBreaks(file.Data)
	if err != nil {
		return nil, err
	}

	var values yaml.Node
	err = yaml.Unmarshal(content, &values)
	if err != nil {
		return nil, fmt.Errorf("error unmarshal values file: %w", err)
	}

	// Variables.
	//for i, v := range sortVariables(mod.Variables) {
	//	// Parse tf expression from type.
	//	var (
	//		tfType       = cty.DynamicPseudoType
	//		defValue     = v.Default
	//		defObj       any
	//		order        = i + 1
	//		tyExpression any
	//	)
	//
	//	// Required and keys.
	//	if v.Required {
	//		required = append(required, v.Name)
	//	}
	//
	//	keys[i] = v.Name
	//
	//	// Generate json schema from tf type or default value.
	//	if v.Type != "" {
	//		// Type exists.
	//		expr, diags := hclsyntax.ParseExpression(stringx.ToBytes(&v.Type), "", hcl.Pos{Line: 1, Column: 1})
	//		if diags.HasErrors() {
	//			return nil, fmt.Errorf("error parsing expression: %w", diags)
	//		}
	//
	//		tfType, defObj, diags = typeexpr.TypeConstraintWithDefaults(expr)
	//		if diags.HasErrors() {
	//			return nil, fmt.Errorf("error getting type: %w", diags)
	//		}
	//
	//		tyExpression = expr
	//	} else if v.Default != nil {
	//		// Empty type, use default value to get type.
	//		b, err := json.Marshal(v.Default)
	//		if err != nil {
	//			return nil, fmt.Errorf("error while marshal terraform variable default value: %w", err)
	//		}
	//
	//		var sjv ctyjson.SimpleJSONValue
	//
	//		err = sjv.UnmarshalJSON(b)
	//		if err != nil {
	//			return nil, fmt.Errorf("error while unmarshal terraform variable default value: %w", err)
	//		}
	//		tfType = sjv.Type()
	//	}
	//
	//	varSchemas.WithProperty(
	//		v.Name,
	//		l.translator.SchemaOfOriginalType(
	//			tfType,
	//			translator.Options{
	//				Name:          v.Name,
	//				DefaultValue:  defValue,
	//				DefaultObject: defObj,
	//				Description:   v.Description,
	//				Sensitive:     v.Sensitive,
	//				Order:         order,
	//				TypeExpress:   tyExpression,
	//			}))
	//}
	//
	// Inject extension sequence.
	sort.Strings(required)
	varSchemas.Required = required
	varSchemas.Extensions = openapi.NewExtFromMap(varSchemas.Extensions).
		WithOriginalVariablesSequence(keys).
		Export()

	// Inject extensions.
	l.injectExts(varSchemas)

	return varSchemas, nil
}

// loadData loads the internal template version data.
func (l *Loader) loadData(chart *chart.Chart) (
	data mytypes.TemplateData, err error,
) {
	// Readme.
	data.Readme = l.getReadme(chart)
	if err != nil {
		return
	}

	// Metadata.
	data.Helm = mytypes.HelmMetadata{
		Chart: *chart.Metadata,
	}

	return
}

// getReadme gets the readme content.
func (l *Loader) getReadme(ct *chart.Chart) string {
	file := findFile(ct, readmeFileNames)
	if file == nil {
		return ""
	}
	return string(file.Data)
}

// injectExts injects extension for variables.
func (l *Loader) injectExts(vs *openapi3.Schema) {
	if vs == nil {
		return
	}

	groupOrder := make(map[string]int)

	for n, v := range vs.Properties {
		if v.Value == nil || v.Value.IsEmpty() {
			continue
		}

		// Group.
		extUI := openapi.GetExtUI(v.Value.Extensions)
		if extUI.Group == "" {
			extUI.Group = defaultGroup

			vs.Properties[n].Value.Extensions = openapi.NewExtFromMap(vs.Properties[n].Value.Extensions).
				WithUIGroup(defaultGroup).
				Export()
		}

		od, ok := groupOrder[extUI.Group]
		if !ok || od > extUI.Order {
			groupOrder[extUI.Group] = extUI.Order
		}
	}

	vsExtUI := openapi.GetExtUI(vs.Extensions)
	if len(vsExtUI.GroupOrder) == 0 {
		ep := openapi.NewExtFromMap(vs.Extensions).
			WithUIGroupOrder(sortMapValue(groupOrder)...).
			Export()
		vs.Extensions = ep
	}
}

func sortMapValue(m map[string]int) []string {
	type keyValue struct {
		Key   string
		Value int
	}

	s := make([]keyValue, 0)

	for key, value := range m {
		s = append(s, keyValue{Key: key, Value: value})
	}

	sort.Slice(s, func(i, j int) bool {
		return s[i].Value < s[j].Value
	})

	keys := make([]string, len(s))
	for i, kv := range s {
		keys[i] = kv.Key
	}

	return keys
}

func sortVariables(m map[string]*tfconfig.Variable) (s []*tfconfig.Variable) {
	s = make([]*tfconfig.Variable, 0, len(m))
	for k := range m {
		s = append(s, m[k])
	}

	sort.SliceStable(s, func(i, j int) bool {
		return judgeSourcePos(&s[i].Pos, &s[j].Pos)
	})

	return
}

func judgeSourcePos(i, j *tfconfig.SourcePos) bool {
	switch {
	case i.Filename < j.Filename:
		return false
	case i.Filename > j.Filename:
		return true
	}

	return i.Line < j.Line
}

func findFile(ct *chart.Chart, fileNames []string) *chart.File {
	for _, f := range ct.Raw {
		for _, n := range fileNames {
			if f == nil {
				continue
			}
			if strings.EqualFold(f.Name, n) {
				return f
			}
		}
	}

	return nil
}

// FormatLineBreaks replaces \r\n with \n in the content.
// When parsing YAML into a yaml.Node, it's advisable to replace \r\n with \n beforehand,
// since YAML content may contain Windows-style line breaks (\r\n).
// This ensures compatibility with YAML parsers, as they typically expect Unix-style line breaks (\n).
func FormatLineBreaks(content []byte) ([]byte, error) {
	return []byte(strings.Replace(string(content), "\r\n", "\n", -1)), nil
}
