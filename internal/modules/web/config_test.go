package web

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseConfig(t *testing.T) {
	tests := []struct {
		name             string
		setup            func()
		teardown         func()
		expectedWebDir   string
		expectedWebTitle string
	}{
		{
			name:             "WEB_DIR environment variable set",
			expectedWebDir:   "/custom_web",
			expectedWebTitle: "test",
			setup: func() {
				os.Setenv("WEB_DIR", "/custom_web")
				os.Setenv("WEB_TITLE", "test")
			},
			teardown: func() {
				os.Unsetenv("WEB_DIR")
				os.Unsetenv("WEB_TITLE")
			},
		},
		{
			name:             "environment variable not set",
			expectedWebDir:   "/web",
			expectedWebTitle: "lazts",
			setup: func() {
				os.Unsetenv("WEB_DIR")
				os.Unsetenv("WEB_TITLE")
			},
			teardown: func() {},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			defer tt.teardown()

			cfg := parseConfig()

			assert.Equal(t, tt.expectedWebDir, cfg.Dir, "WebDir should be equal")
			assert.Equal(t, tt.expectedWebTitle, cfg.Title, "WebTitle should be equal")
		})
	}
}
