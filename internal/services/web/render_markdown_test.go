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
	mock := new(markdown.Mock)

	os.Setenv("WEB_DIR", testDir)
	os.Setenv("WEB_TITLE", testTitle)
	defer os.Unsetenv("WEB_DIR")
	defer os.Unsetenv("WEB_TITLE")
	defer utils.RemoveTestDir(t, testDir)

	utils.CreateTestFile(t, filepath.Join(testDir, "templates", "layouts"), "base.html", `{{define "base"}}<html><head></head><body>{{template "content" .}}</body></html>{{end}}`)
	utils.CreateTestFile(t, filepath.Join(testDir, "templates", "pages"), "blog-content.html", `<h2>Blog</h2><article>{{template "markdown" .}}</article>`)

	r := New(mock)

	tests := []struct {
		name      string
		path      string
		content   string
		setup     func()
		want      string
		expectErr bool
	}{
		{
			name:    "Render template",
			path:    "blog-content",
			content: "sample",
			setup: func() {
				mock.On("LoadContent", "blog", "sample").Return(`<h1>Blog Page 1</h1>`, nil).Once()
				mock.On("LoadMetadata", "blog", "sample").Return(map[string]interface{}{}, nil).Once()
			},
			want:      `<html><head></head><body><h2>Blog</h2><article><h1>Blog Page 1</h1></article></body></html>`,
			expectErr: false,
		},
		{
			name:    "Render template with error",
			path:    "blog-content",
			content: "sample",
			setup: func() {
				mock.On("LoadContent", "blog", "sample").Return("", assert.AnError).Once()
			},
			expectErr: true,
		},
		{
			name:    "Render template with no content",
			path:    "blog-content",
			content: "sample",
			setup: func() {
				mock.On("LoadContent", "blog", "sample").Return("", nil).Once()
				mock.On("LoadMetadata", "blog", "sample").Return(map[string]interface{}{}, nil).Once()
			},
			want:      `<html><head></head><body><h2>Blog</h2><article></article></body></html>`,
			expectErr: false,
		},
		{
			name:    "Render not found template",
			path:    "blog-content",
			content: "nonexist",
			setup: func() {
				mock.On("LoadContent", "blog", "nonexist").Return("", assert.AnError).Once()
			},
			expectErr: true,
		},
		{
			name:    "Render with invalid path",
			path:    "",
			content: "sample",
			setup: func() {
			},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			tt.setup()
			err := r.RenderMarkdown(&buf, tt.path, tt.content, map[string]interface{}{})
			if tt.expectErr {
				assert.Error(t, err, "expected an error but got none")
			} else {
				assert.NoError(t, err, "unexpected error during rendering")
			}

			assert.Equal(t, tt.want, buf.String(), "unexpected output from rendering")
			mock.AssertExpectations(t)
		})
	}
}
