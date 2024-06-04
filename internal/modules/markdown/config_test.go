package markdown

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseConfig(t *testing.T) {
	tests := []struct {
		name          string
		setup         func()
		teardown      func()
		expectedTheme string
	}{
		{
			name:          "environment variable set",
			expectedTheme: "test",
			setup: func() {
				os.Setenv("MARKDOWN_THEME", "test")
			},
			teardown: func() {
				os.Unsetenv("MARKDOWN_THEME")
			},
		},
		{
			name:          "environment variable not set",
			expectedTheme: DEFAULT_THEME,
			setup: func() {
				os.Unsetenv("MARKDOWN_THEME")
			},
			teardown: func() {},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			defer tt.teardown()

			cfg := parseConfig()

			assert.Equal(t, tt.expectedTheme, cfg.Theme, "theme should be equal")
		})
	}
}
