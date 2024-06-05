package web

import (
	"bytes"
	"lazts/internal/utils"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRenderPartial(t *testing.T) {
	testDir := "test_data"
	os.Setenv("WEB_DIR", testDir)
	defer os.Unsetenv("WEB_DIR")
	defer utils.RemoveTestDir(t, testDir)

	utils.CreateTestFile(t, filepath.Join(testDir, "templates", "layouts"), "base.html", `{{define "base"}}<html><body>{{template "partials" .}}</body></html>{{end}}`)
	utils.CreateTestFile(t, filepath.Join(testDir, "templates", "partials"), "partial.html", `{{define "partial"}}Partial template: {{.Data}}{{end}}`)

	r := New(nil)

	tests := []struct {
		name      string
		partial   string
		data      map[string]interface{}
		want      string
		expectErr bool
	}{
		{
			name:      "Successful rendering",
			partial:   "partial",
			data:      map[string]interface{}{"Data": "test"},
			want:      "Partial template: test",
			expectErr: false,
		},
		{
			name:      "Failed to clone templates",
			partial:   "partial",
			data:      nil,
			want:      "Partial template: <no value>",
			expectErr: false,
		},
		{
			name:      "Failed to execute template",
			partial:   "nonexistent",
			data:      map[string]interface{}{"Data": "test"},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := r.RenderPartial(&buf, tt.partial, tt.data)
			if tt.expectErr {
				assert.Error(t, err, "expected an error but got none")
			} else {
				assert.NoError(t, err, "unexpected error during rendering")
				assert.Equal(t, tt.want, buf.String(), "unexpected output from rendering")
			}
		})
	}
}
