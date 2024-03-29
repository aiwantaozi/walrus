package helm

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/yaml"
)

func TestLoadTerraformSchema(t *testing.T) {
	testCases := []struct {
		name          string
		input         string
		expectedError bool
	}{
		//{
		//	name:          "Invalid",
		//	input:         "testdata/invalid",
		//	expectedError: true,
		//},
		{
			name:          "With README.md",
			input:         "testdata/with_readme",
			expectedError: false,
		},
		//{
		//	name:          "With variable",
		//	input:         "testdata/with_variable",
		//	expectedError: false,
		//},
		//{
		//	name:          "With description",
		//	input:         "testdata/with_description",
		//	expectedError: false,
		//},
		//{
		//	name:          "With schema.yaml",
		//	input:         "testdata/full_schema",
		//	expectedError: false,
		//},
		//{
		//	name:          "Complex variable",
		//	input:         "testdata/complex_variable",
		//	expectedError: false,
		//},
		//{
		//	name:          "Complex default",
		//	input:         "testdata/complex_default",
		//	expectedError: false,
		//},
		//{
		//	name:          "With any variable",
		//	input:         "testdata/any_variable",
		//	expectedError: false,
		//},
		//{
		//	name:          "With only default",
		//	input:         "testdata/with_only_default",
		//	expectedError: false,
		//},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			loader := New()

			actualOutput, actualError := loader.Load(tc.input, "dev-template")
			assert.Equal(t, tc.expectedError, actualError != nil)

			if tc.expectedError {
				return
			}

			b, err := json.Marshal(actualOutput)
			assert.NoError(t, err)
			actual, err := yaml.JSONToYAML(b)
			assert.NoError(t, err)

			expected, err := os.ReadFile(filepath.Join(tc.input, "expected.yaml"))
			assert.NoError(t, err)

			equal := assert.Equal(t, expected, actual)
			if !equal {
				_ = os.WriteFile(filepath.Join(tc.input, "actual.yaml"), actual, 0o644)
			}
		})
	}
}
