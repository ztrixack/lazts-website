package watermark

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseConfig(t *testing.T) {
	tests := []struct {
		name         string
		envVars      map[string]string
		expectedDir  string
		expectedPath string
		expectedSize int
	}{
		{
			name:         "Default values",
			envVars:      map[string]string{},
			expectedDir:  DEFAULT_DIR,
			expectedPath: DEFAULT_PATH,
			expectedSize: DEFAULT_SIZE,
		},
		{
			name:         "Custom WEB_DIR",
			envVars:      map[string]string{"WEB_DIR": "/custom/dir"},
			expectedDir:  "/custom/dir",
			expectedPath: DEFAULT_PATH,
			expectedSize: DEFAULT_SIZE,
		},
		{
			name:         "Custom WATERMARK_PATH",
			envVars:      map[string]string{"WATERMARK_PATH": "/custom/path"},
			expectedDir:  DEFAULT_DIR,
			expectedPath: "/custom/path",
			expectedSize: DEFAULT_SIZE,
		},
		{
			name:         "Custom WATERMARK_SIZE",
			envVars:      map[string]string{"WATERMARK_SIZE": "100"},
			expectedDir:  DEFAULT_DIR,
			expectedPath: DEFAULT_PATH,
			expectedSize: 100,
		},
		{
			name:         "Invalid WATERMARK_SIZE",
			envVars:      map[string]string{"WATERMARK_SIZE": "invalid"},
			expectedDir:  DEFAULT_DIR,
			expectedPath: DEFAULT_PATH,
			expectedSize: DEFAULT_SIZE,
		},
		{
			name: "All custom values",
			envVars: map[string]string{
				"WEB_DIR":        "/custom/dir",
				"WATERMARK_PATH": "/custom/path",
				"WATERMARK_SIZE": "100",
			},
			expectedDir:  "/custom/dir",
			expectedPath: "/custom/path",
			expectedSize: 100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set environment variables
			for key, value := range tt.envVars {
				os.Setenv(key, value)
			}

			// Parse config
			cfg := parseConfig()

			// Assertions
			assert.Equal(t, tt.expectedDir, cfg.Dir, "Dir should match expected value")
			assert.Equal(t, tt.expectedPath, cfg.Path, "Path should match expected value")
			assert.Equal(t, tt.expectedSize, cfg.Size, "Size should match expected value")

			// Unset environment variables
			for key := range tt.envVars {
				os.Unsetenv(key)
			}
		})
	}
}
