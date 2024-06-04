package web

import (
	"bytes"
	"lazts/internal/modules/markdown"
	"lazts/internal/utils"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWeb_RenderMarkdown(t *testing.T) {
	testDir := "test_data"
	testTitle := "test_title"
	markdownMock := new(markdown.Mock)

	os.Setenv("WEB_DIR", testDir)
	os.Setenv("WEB_TITLE", testTitle)
	defer os.Unsetenv("WEB_DIR")
	defer os.Unsetenv("WEB_TITLE")
	defer utils.RemoveTestDir(t, testDir)

	utils.CreateTestFile(t, filepath.Join(testDir, "templates", "layouts"), "base.html", `{{define "base"}}<html><head><title>{{.Title}}</title></head><body>{{template "content" .}}</body></html>{{end}}`)
	utils.CreateTestFile(t, filepath.Join(testDir, "templates", "pages"), "blog_content.html", `<h2>Blog</h2><article>{{template "markdown" .}}</article>`)

	r := New(markdownMock).(*service)

	tests := []struct {
		name      string
		path      string
		setup     func()
		want      string
		expectErr bool
	}{
		{
			name: "Render template",
			path: "/blog/sample",
			setup: func() {
				markdownMock.On("ToHTML", filepath.Join(testDir, "contents", "blog", "sample", "page.md")).Return(`<h1>Blog Page 1</h1>`, nil).Once()
			},
			want:      `<html><head><title>test_title</title></head><body><h2>Blog</h2><article><h1>Blog Page 1</h1></article></body></html>`,
			expectErr: false,
		},
		{
			name: "Render template with error",
			path: "/blog/sample",
			setup: func() {
				markdownMock.On("ToHTML", filepath.Join(testDir, "contents", "blog", "sample", "page.md")).Return("", assert.AnError).Once()
			},
			expectErr: true,
		},
		{
			name: "Render template with no content",
			path: "/blog/sample",
			setup: func() {
				markdownMock.On("ToHTML", filepath.Join(testDir, "contents", "blog", "sample", "page.md")).Return("", nil).Once()
			},
			want:      `<html><head><title>test_title</title></head><body><h2>Blog</h2><article></article></body></html>`,
			expectErr: false,
		},
		{
			name: "Render not found template",
			path: "/blog/nonexist",
			setup: func() {
				markdownMock.On("ToHTML", filepath.Join(testDir, "contents", "blog", "nonexist", "page.md")).Return("", assert.AnError).Once()
			},
			expectErr: true,
		},
		{
			name: "Render with invalid path",
			path: "/blog/../sample",
			setup: func() {
				markdownMock.On("ToHTML", filepath.Join(testDir, "contents", "sample", "page.md")).Return("", assert.AnError).Once()
			},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			tt.setup()
			err := r.RenderMarkdown(&buf, tt.path)
			if tt.expectErr {
				assert.Error(t, err, "expected an error but got none")
			} else {
				assert.NoError(t, err, "unexpected error during rendering")
			}

			assert.Equal(t, tt.want, buf.String(), "unexpected output from rendering")
			markdownMock.AssertExpectations(t)
		})
	}
}
