package memo

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseConfig(t *testing.T) {
	tests := []struct {
		name               string
		setup              func()
		teardown           func()
		expectedContentDir string
	}{
		{
			name:               "environment variable set",
			expectedContentDir: "/custom_content",
			setup: func() {
				os.Setenv("CONTENT_DIR", "/custom_content")
			},
			teardown: func() {
				os.Unsetenv("CONTENT_DIR")
			},
		},
		{
			name:               "environment variable not set",
			expectedContentDir: DEFAULT_CONTENT_DIR,
			setup: func() {
				os.Unsetenv("CONTENT_DIR")
			},
			teardown: func() {},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			defer tt.teardown()

			cfg := parseConfig()

			assert.Equal(t, tt.expectedContentDir, cfg.ContentDir, "ContentDir should be equal")
		})
	}
}
