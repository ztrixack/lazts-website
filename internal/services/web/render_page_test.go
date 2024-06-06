package web

import (
	"bytes"
	"lazts/internal/utils"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRenderPage(t *testing.T) {
	const WEB_DIR = "test_render_page"
	const WEB_TITLE = "Test Render Page"

	tests := []struct {
		name           string
		path           string
		expectedError  bool
		expectedResult string
		setup          func(t *testing.T)
	}{
		{
			name:           "Render template",
			path:           "blog",
			expectedResult: `<html><head></head><body><h2>Blog</h2><p>This is the blog page.</p></body></html>`,
			expectedError:  false,
			setup: func(t *testing.T) {
				utils.CreateTestFile(t, filepath.Join(WEB_DIR, "templates", "layouts"), "base.html", `{{define "base"}}<html><head></head><body>{{template "content" .}}</body></html>{{end}}`)
				utils.CreateTestFile(t, filepath.Join(WEB_DIR, "templates", "pages"), "home.html", `<h2>Home</h2><p>This is the home page.</p>`)
				utils.CreateTestFile(t, filepath.Join(WEB_DIR, "templates", "pages"), "blog.html", `<h2>Blog</h2><p>This is the blog page.</p>`)
			},
		},
		{
			name:           "Render home template",
			path:           "home",
			expectedResult: `<html><head></head><body><h2>Home</h2><p>This is the home page.</p></body></html>`,
			expectedError:  false,
			setup: func(t *testing.T) {
				utils.CreateTestFile(t, filepath.Join(WEB_DIR, "templates", "layouts"), "base.html", `{{define "base"}}<html><head></head><body>{{template "content" .}}</body></html>{{end}}`)
				utils.CreateTestFile(t, filepath.Join(WEB_DIR, "templates", "pages"), "home.html", `<h2>Home</h2><p>This is the home page.</p>`)
				utils.CreateTestFile(t, filepath.Join(WEB_DIR, "templates", "pages"), "blog.html", `<h2>Blog</h2><p>This is the blog page.</p>`)
			},
		},
		{
			name:          "No render with empty template",
			path:          "",
			expectedError: true,
			setup:         func(t *testing.T) {},
		},
		{
			name:          "No render not found template",
			path:          "nonexist",
			expectedError: true,
			setup:         func(t *testing.T) {},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup(t)
			r, m := setup(WEB_DIR, WEB_TITLE)
			defer teardown(t, WEB_DIR)

			var buf bytes.Buffer
			err := r.RenderPage(&buf, tt.path, nil)

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
