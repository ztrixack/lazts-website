package web

import (
	"bytes"
	"lazts/internal/modules/markdown"
	"lazts/internal/utils"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRenderMarkdown(t *testing.T) {
	const WEB_DIR = "test_render_partial"
	const WEB_TITLE = "Test Render Partial"

	tests := []struct {
		name           string
		path           string
		content        string
		expectedError  bool
		expectedResult string
		setup          func(t *testing.T)
		setupMock      func(mock *markdown.Mock)
	}{
		{
			name:    "Render template",
			path:    "blog-content",
			content: "sample",
			setup: func(t *testing.T) {
				utils.CreateTestFile(t, filepath.Join(WEB_DIR, "templates", "layouts"), "base.html", `{{define "base"}}<html><head></head><body>{{template "content" .}}</body></html>{{end}}`)
				utils.CreateTestFile(t, filepath.Join(WEB_DIR, "templates", "pages"), "blog-content.html", `<h2>Blog</h2><article>{{template "markdown" .}}</article>`)
			},
			setupMock: func(mock *markdown.Mock) {
				mock.On("LoadContent", "blog", "sample").Return(`<h1>Blog Page 1</h1>`, nil).Once()
				mock.On("LoadMetadata", "blog", "sample").Return(map[string]interface{}{}, nil).Once()
			},
			expectedResult: `<html><head></head><body><h2>Blog</h2><article><h1>Blog Page 1</h1></article></body></html>`,
			expectedError:  false,
		},
		{
			name:    "Render template with error",
			path:    "blog-content",
			content: "sample",
			setup: func(t *testing.T) {
				utils.CreateTestFile(t, filepath.Join(WEB_DIR, "templates", "layouts"), "base.html", `{{define "base"}}<html><head></head><body>{{template "content" .}}</body></html>{{end}}`)
				utils.CreateTestFile(t, filepath.Join(WEB_DIR, "templates", "pages"), "blog-content.html", `<h2>Blog</h2><article>{{template "markdown" .}}</article>`)
			},
			setupMock: func(mock *markdown.Mock) {
				mock.On("LoadContent", "blog", "sample").Return("", assert.AnError).Once()
			},
			expectedError: true,
		},
		{
			name:           "Render template with no content",
			path:           "blog-content",
			content:        "sample",
			expectedError:  false,
			expectedResult: `<html><head></head><body><h2>Blog</h2><article></article></body></html>`,
			setup: func(t *testing.T) {
				utils.CreateTestFile(t, filepath.Join(WEB_DIR, "templates", "layouts"), "base.html", `{{define "base"}}<html><head></head><body>{{template "content" .}}</body></html>{{end}}`)
				utils.CreateTestFile(t, filepath.Join(WEB_DIR, "templates", "pages"), "blog-content.html", `<h2>Blog</h2><article>{{template "markdown" .}}</article>`)
			},
			setupMock: func(mock *markdown.Mock) {
				mock.On("LoadContent", "blog", "sample").Return("", nil).Once()
				mock.On("LoadMetadata", "blog", "sample").Return(map[string]interface{}{}, nil).Once()
			},
		},
		{
			name:          "Render not found template",
			path:          "blog-content",
			content:       "nonexist",
			expectedError: true,
			setup:         func(t *testing.T) {
				utils.CreateTestFile(t, filepath.Join(WEB_DIR, "templates", "layouts"), "base.html", `{{define "base"}}<html><head></head><body>{{template "content" .}}</body></html>{{end}}`)
				utils.CreateTestFile(t, filepath.Join(WEB_DIR, "templates", "pages"), "blog-content.html", `<h2>Blog</h2><article>{{template "markdown" .}}</article>`)
			},
			setupMock: func(mock *markdown.Mock) {
				mock.On("LoadContent", "blog", "nonexist").Return("", assert.AnError).Once()
			},
		},
		{
			name:          "Render with invalid path",
			path:          "",
			content:       "sample",
			expectedError: true,
			setup:         func(t *testing.T) {},
			setupMock:     func(mock *markdown.Mock) {},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup(t)
			r, m := setup(WEB_DIR, WEB_TITLE)
			tt.setupMock(m)
			defer teardown(t, WEB_DIR)

			var buf bytes.Buffer
			err := r.RenderMarkdown(&buf, tt.path, tt.content, map[string]interface{}{})
			if tt.expectedError {
				assert.Error(t, err, "expected an error but got none")
			} else {
				assert.NoError(t, err, "unexpected error during rendering")
				assert.Equal(t, tt.expectedResult, buf.String(), "unexpected output from rendering")
			}

			m.AssertExpectations(t)
		})
	}
}
