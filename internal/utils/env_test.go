package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLookupEnv(t *testing.T) {
	tests := []struct {
		key      string
		value    string
		expected string
		setup    func()
		teardown func()
	}{
		{
			key:      "EXISTING_KEY",
			value:    "existing_value",
			expected: "existing_value",
			setup: func() {
				os.Setenv("EXISTING_KEY", "existing_value")
			},
			teardown: func() {
				os.Unsetenv("EXISTING_KEY")
			},
		},
		{
			key:      "NON_EXISTING_KEY",
			value:    "",
			expected: "default_value",
			setup:    func() {},
			teardown: func() {},
		},
	}

	for _, tt := range tests {
		tt.setup()
		actual := LookupEnv(tt.key, "default_value")
		assert.Equal(t, tt.expected, actual, "they should be equal")
		tt.teardown()
	}
}
