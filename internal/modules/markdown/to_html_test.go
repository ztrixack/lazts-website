package markdown

import (
	"lazts/internal/utils"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToHTML(t *testing.T) {
	testDir := "test_data"
	m := New().(*module)

	utils.CreateTestFile(t, testDir, "test.md", "# Hello, World!\n\nThis is a test.")
	defer utils.RemoveTestDir(t, testDir)

	tests := []struct {
		name            string
		path            string
		setup           func()
		expectedContent string
		expectedError   bool
	}{
		{
			name:            "successful markdown conversion",
			path:            filepath.Join(testDir, "test.md"),
			setup:           func() {},
			expectedContent: "<h1 id=\"hello-world\">Hello, World!</h1>",
		},
		{
			name: "cache hit",
			path: "test.md",
			setup: func() {
				m.cache["test.md"] = "<p>Cached content</p>"
			},
			expectedContent: "<p>Cached content</p>",
		},
		{
			name:          "file read error",
			path:          "nonexistent.md",
			setup:         func() {},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()

			content, err := m.ToHTML(tt.path)
			if tt.expectedError {
				assert.Error(t, err, "expected error")
				assert.Equal(t, "", content, "expected empty content on error")
			} else {
				assert.NoError(t, err, "expected no error")
				assert.Contains(t, content, tt.expectedContent, "expected content mismatch")
			}
		})
	}
}
