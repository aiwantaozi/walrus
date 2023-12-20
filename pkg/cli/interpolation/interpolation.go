package interpolation

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/compose-spec/compose-go/interpolation"
	"github.com/compose-spec/compose-go/template"
)

const (
	APIGroupVariables = "variables"
)

var (
	delimiter            = "\\$"
	substitutionNamed    = "[_a-z][_a-z0-9]*"
	substitutionBraced   = "[_a-z][_a-z0-9]*(?::?[-+?](.*))?"
	substitutionVariable = "{var\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?}"
	substitutionResource = "{res\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?\\.[_a-z][_a-z0-9]*}"
	substitutionService  = "{svc\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?\\.[_a-z][_a-z0-9]*}"
)

var (
	groupEscaped  = "escaped"
	groupNamed    = "named"
	groupBraced   = "braced"
	groupVariable = "variable"
	groupResource = "resource"
	groupService  = "service"
	groupInvalid  = "invalid"
)

// patternString is adapted from:
// https://github.com/compose-spec/compose-go/blob/81e1e9036e66e9afcdbecffea6470ff7edcffef8/template/template.go#L38
var patternString = fmt.Sprintf(
	"%s(?i:(?P<%s>%s)|(?P<%s>%s)|(?P<%s>%s)|(?P<%s>%s)|(?P<%s>%s)|{(?:(?P<%s>%s)}|(?P<%s>)))",
	delimiter,
	groupEscaped, delimiter,
	groupNamed, substitutionNamed,
	groupVariable, substitutionVariable,
	groupResource, substitutionResource,
	groupService, substitutionService,
	groupBraced, substitutionBraced,
	groupInvalid,
)

var defaultPattern = regexp.MustCompile(patternString)

// variableFilePattern is a regular expression pattern used for identifying file path start with @.
var variableFilePattern = regexp.MustCompile(`@([^@\s]+)`)

// Interpolate interpolates the given yaml object.
func Interpolate(yml map[string]any, data map[string]string, interps ...Interpolator) (map[string]any, error) {
	if len(interps) == 0 {
		interps = DefaultInterpolator()
	}

	var err error
	for _, v := range interps {
		yml, err = v.Func(yml, data)
		if err != nil {
			return nil, fmt.Errorf("failed to interpolate with %s: %w", v.Name, err)
		}
	}

	return yml, nil
}

// DefaultInterpolator returns the default interpolates.
func DefaultInterpolator() []Interpolator {
	return []Interpolator{
		ContextAndEnvironmentVariableInterpolator(),
		VariableFileInterpolator(),
	}
}

// Interpolator represents a type used for interpolating values in a YAML object.
type Interpolator struct {
	Name string
	Func InterpolateFunc
}

// InterpolateFunc is a function that interpolates values in a YAML object.
type InterpolateFunc = func(ymlObj map[string]any, data map[string]string) (map[string]any, error)

// ContextAndEnvironmentVariableInterpolator interpolates the context and environment variable.
func ContextAndEnvironmentVariableInterpolator() Interpolator {
	substitute := func(tmpl string, mapping template.Mapping) (string, error) {
		return template.SubstituteWithOptions(
			tmpl,
			mapping,
			template.WithPattern(defaultPattern),
			template.WithReplacementFunction(func(s string, m template.Mapping, c *template.Config) (string, error) {
				switch {
				case strings.HasPrefix(s, "${var.") ||
					strings.HasPrefix(s, "${svc.") ||
					strings.HasPrefix(s, "${res."):
					return s, nil
				default:
					return template.DefaultReplacementFunc(s, m, c)
				}
			}),
		)
	}

	return Interpolator{
		Name: "context-and-environment-variable-interpolator",
		Func: func(ymlObj map[string]any, data map[string]string) (map[string]any, error) {
			return interpolation.Interpolate(
				ymlObj,
				interpolation.Options{
					LookupValue: func(key string) (string, bool) {
						switch key {
						case "Project":
							return data["project"], true
						case "Environment":
							return data["environment"], true
						default:
							return os.LookupEnv(key)
						}
					},
					Substitute: substitute,
				})
		},
	}
}

// VariableFileInterpolator interpolates variables from files.
func VariableFileInterpolator() Interpolator {
	substitute := func(tmpl string, mapping template.Mapping) (string, error) {
		return template.SubstituteWithOptions(
			tmpl,
			mapping,
			template.WithPattern(variableFilePattern),
			template.WithReplacementFunction(func(s string, m template.Mapping, c *template.Config) (string, error) {
				file := strings.TrimPrefix(s, "@")
				if !filepath.IsAbs(file) {
					dir, err := os.Getwd()
					if err != nil {
						return "", err
					}

					file = filepath.Join(dir, file)
				}

				b, err := os.ReadFile(file)
				if err != nil {
					return "", err
				}

				return string(b), nil
			}),
		)
	}

	return Interpolator{
		Name: "variable-file-interpolator",
		Func: func(ymlObj map[string]any, _ map[string]string) (map[string]any, error) {
			opts := interpolation.Options{
				Substitute: substitute,
			}

			var (
				variables = make(map[string]any)
				ok        bool
			)

			variables[APIGroupVariables], ok = ymlObj[APIGroupVariables]
			if !ok {
				return ymlObj, nil
			}

			interp, err := interpolation.Interpolate(variables, opts)
			if err != nil {
				return nil, err
			}

			ymlObj[APIGroupVariables] = interp[APIGroupVariables]
			if err != nil {
				return nil, err
			}

			return ymlObj, nil
		},
	}
}
