package web

import (
	"bytes"
	"lazts/internal/utils"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRenderPartial(t *testing.T) {
	const WEB_DIR = "test_render_partial"
	const WEB_TITLE = "Test Render Partial"

	tests := []struct {
		name           string
		partial        string
		data           map[string]interface{}
		expectedError  bool
		expectedResult string
		setup          func(t *testing.T)
	}{
		{
			name:           "Successful rendering",
			partial:        "partial",
			data:           map[string]interface{}{"Data": "test"},
			expectedError:  false,
			expectedResult: "Partial template: test",
			setup: func(t *testing.T) {
				utils.CreateTestFile(t, filepath.Join(WEB_DIR, "templates", "layouts"), "base.html", `{{define "base"}}<html><body>{{template "partials" .}}</body></html>{{end}}`)
				utils.CreateTestFile(t, filepath.Join(WEB_DIR, "templates", "partials"), "partial.html", `{{define "partial"}}Partial template: {{.Data}}{{end}}`)
			},
		},
		{
			name:           "Failed to clone templates",
			partial:        "partial",
			data:           nil,
			expectedError:  false,
			expectedResult: "Partial template: <no value>",
			setup: func(t *testing.T) {
				utils.CreateTestFile(t, filepath.Join(WEB_DIR, "templates", "layouts"), "base.html", `{{define "base"}}<html><body>{{template "partials" .}}</body></html>{{end}}`)
				utils.CreateTestFile(t, filepath.Join(WEB_DIR, "templates", "partials"), "partial.html", `{{define "partial"}}Partial template: {{.Data}}{{end}}`)
			},
		},
		{
			name:          "Failed to execute template",
			partial:       "nonexistent",
			data:          map[string]interface{}{"Data": "test"},
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
			err := r.RenderPartial(&buf, tt.partial, tt.data)

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
