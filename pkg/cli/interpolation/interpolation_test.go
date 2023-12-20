package interpolation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInterpolation(t *testing.T) {
	t.Setenv("FOO", "foo")
	t.Setenv("FOO2", "foo2")

	testcases := []struct {
		test     string
		expected string
		isErr    bool
	}{
		{test: "${FOO}", expected: "foo"},
		{test: "${FOO-foo}", expected: "foo"},
		{test: "${FOO:-foo_}", expected: "foo"},
		{test: "${BAR:-bar}", expected: "bar"},

		{test: "${FOO2:?error message}", expected: "foo2"},
		{test: "${FOO2?error message}", expected: "foo2"},
		{test: "${BAR:?error message}", isErr: true},
		{test: "${BAR?error message}", isErr: true},

		// Built-in.
		{test: "${var.test}", expected: "${var.test}"},
		{test: "${resource.test.output}", expected: "${resource.test.output}"},
		{test: "${service.test.output}", expected: "${service.test.output}"},
		{test: "${unsupported.test}", isErr: true},

		// Variable file.
		{test: "@testdata/env", expected: "foo"},
	}

	for _, tc := range testcases {
		yml := map[string]any{
			"variables": tc.test,
		}
		result, err := Interpolate(yml, nil)
		assert.Equal(t, err != nil, tc.isErr)

		if err == nil {
			assert.Equal(t, tc.expected, result["variables"])
		}
	}
}
