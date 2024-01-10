package openapi

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/stretchr/testify/assert"
)

func TestTraverse(t *testing.T) {
	testCases := []struct {
		name          string
		input         string
		expectedError bool
	}{
		{
			name:          "Empty",
			input:         "testdata/empty",
			expectedError: true,
		},
		{
			name:          "With simple",
			input:         "testdata/simple",
			expectedError: true,
		},
		{
			name:          "With map",
			input:         "testdata/map",
			expectedError: false,
		},
		{
			name:          "With list",
			input:         "testdata/list",
			expectedError: false,
		},
		{
			name:          "With object",
			input:         "testdata/object",
			expectedError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l := openapi3.NewLoader()

			ts, err := l.LoadFromFile(filepath.Join(tc.input, "schema.yaml"))
			assert.NoError(t, err)

			s := ts.Components.Schemas["variables"]
			m, err := GenSchemaDefaultPatch(context.Background(), s.Value)
			assert.NoError(t, err)

			eb, err := os.ReadFile(filepath.Join(tc.input, "expected.json"))
			assert.NoError(t, err)

			if len(eb) == 0 {
				assert.Empty(t, m)
				return
			}

			assert.Equal(t, eb, m)
		})
	}
}
