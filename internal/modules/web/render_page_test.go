package web

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWeb_RenderPage(t *testing.T) {
	testDir := "test_data"
	testTitle := "test_title"
	os.Setenv("WEB_DIR", testDir)
	os.Setenv("WEB_TITLE", testTitle)
	defer os.Unsetenv("WEB_DIR")
	defer os.Unsetenv("WEB_TITLE")
	defer removeTestDir(t, testDir)

	createTestFile(t, filepath.Join(testDir, "templates", "layouts"), "base.html", `{{define "base"}}<html><head><title>{{.Title}}</title></head><body>{{template "content" .}}</body></html>{{end}}`)
	createTestFile(t, filepath.Join(testDir, "templates", "pages"), "home.html", `<h2>Home</h2><p>This is the home page.</p>`)
	createTestFile(t, filepath.Join(testDir, "templates", "pages"), "blog.html", `<h2>Blog</h2><p>This is the blog page.</p>`)

	r := New().(*webber)

	tests := []struct {
		name      string
		path      string
		want      string
		expectErr bool
	}{
		{
			name:      "Render template",
			path:      "/blog",
			want:      `<html><head><title>test_title</title></head><body><h2>Blog</h2><p>This is the blog page.</p></body></html>`,
			expectErr: false,
		},
		{
			name:      "Render template without /",
			path:      "blog",
			want:      `<html><head><title>test_title</title></head><body><h2>Blog</h2><p>This is the blog page.</p></body></html>`,
			expectErr: false,
		},
		{
			name:      "Render home template",
			path:      "/",
			want:      `<html><head><title>test_title</title></head><body><h2>Home</h2><p>This is the home page.</p></body></html>`,
			expectErr: false,
		},
		{
			name:      "Render not found template",
			path:      "nonexist",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := r.RenderPage(&buf, tt.path)
			if tt.expectErr {
				assert.Error(t, err, "expected an error but got none")
			} else {
				assert.NoError(t, err, "unexpected error during rendering")
			}

			assert.Equal(t, tt.want, buf.String(), "unexpected output from rendering")
		})
	}
}

func createTestFile(t *testing.T, dir, name, content string) {
	t.Helper()
	filePath := filepath.Join(dir, name)
	err := os.MkdirAll(filepath.Dir(filePath), 0755)
	if err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}
	err = os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
}

func removeTestDir(t *testing.T, dir string) {
	t.Helper()
	err := os.RemoveAll(dir)
	if err != nil {
		t.Fatalf("Failed to remove test directory: %v", err)
	}
}
