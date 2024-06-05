package web

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseConfig(t *testing.T) {
	tests := []struct {
		name            string
		setup           func()
		teardown        func()
		expectedDir     string
		expectedTitle   string
		expectedExcerpt string
	}{
		{
			name:            "environment variable set",
			expectedDir:     "/custom_web",
			expectedTitle:   "test",
			expectedExcerpt: "excerpt",
			setup: func() {
				os.Setenv("WEB_DIR", "/custom_web")
				os.Setenv("WEB_TITLE", "test")
				os.Setenv("WEB_EXCERPT", "excerpt")
			},
			teardown: func() {
				os.Unsetenv("WEB_DIR")
				os.Unsetenv("WEB_TITLE")
				os.Unsetenv("WEB_EXCERPT")
			},
		},
		{
			name:            "environment variable not set",
			expectedDir:     DEFAULT_DIR,
			expectedTitle:   DEFAULT_TITLE,
			expectedExcerpt: DEFAULT_EXCERPT,
			setup: func() {
				os.Unsetenv("WEB_DIR")
				os.Unsetenv("WEB_TITLE")
				os.Unsetenv("WEB_EXCERPT")
			},
			teardown: func() {},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			defer tt.teardown()

			cfg := parseConfig()

			assert.Equal(t, tt.expectedDir, cfg.Dir, "WebDir should be equal")
			assert.Equal(t, tt.expectedTitle, cfg.Title, "WebTitle should be equal")
			assert.Equal(t, tt.expectedExcerpt, cfg.Excerpt, "WebExcerpt should be equal")
		})
	}
}
