package web

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWeb_RenderMarkdown(t *testing.T) {
	testDir := "test_data"
	testTitle := "test_title"
	os.Setenv("WEB_DIR", testDir)
	os.Setenv("WEB_TITLE", testTitle)
	defer os.Unsetenv("WEB_DIR")
	defer os.Unsetenv("WEB_TITLE")
	defer removeTestDir(t, testDir)

	createTestFile(t, filepath.Join(testDir, "templates", "layouts"), "base.html", `{{define "base"}}<html><head><title>{{.Title}}</title></head><body>{{template "content" .}}</body></html>{{end}}`)
	createTestFile(t, filepath.Join(testDir, "templates", "pages"), "blog_content.html", `<h2>Blog</h2><article>{{template "markdown" .}}</article>`)
	createTestFile(t, filepath.Join(testDir, "contents", "blog", "sample"), "page.md", `# Blog Page 1`)

	r := New().(*webber)

	tests := []struct {
		name      string
		path      string
		want      string
		expectErr bool
	}{
		{
			name:      "Render template",
			path:      "/blog/sample",
			want:      `<html><head><title>test_title</title></head><body><h2>Blog</h2><article><h1>Blog Page 1</h1></article></body></html>`,
			expectErr: false,
		},
		{
			name:      "Render template without /",
			path:      "blog/sample",
			want:      `<html><head><title>test_title</title></head><body><h2>Blog</h2><article><h1>Blog Page 1</h1></article></body></html>`,
			expectErr: false,
		},
		{
			name:      "Render not found template",
			path:      "/blog/nonexist",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := r.RenderMarkdown(&buf, tt.path)
			if tt.expectErr {
				assert.Error(t, err, "expected an error but got none")
			} else {
				assert.NoError(t, err, "unexpected error during rendering")
			}

			assert.Equal(t, tt.want, buf.String(), "unexpected output from rendering")
		})
	}
}
