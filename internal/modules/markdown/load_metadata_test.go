package markdown

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarkdown_LoadMetadata(t *testing.T) {
	m := setup(t)
	defer teardown(t)

	tests := []struct {
		name         string
		domain       string
		slug         string
		expectErr    bool
		expectResult map[string]interface{}
	}{
		{
			name:         "Successful metadata loading",
			domain:       "domain",
			slug:         "slug",
			expectErr:    false,
			expectResult: map[string]interface{}{"meta": "test", "read_time": 0},
		},
		{
			name:         "Invalid domain",
			domain:       "nonexists",
			slug:         "slug",
			expectErr:    true,
			expectResult: nil,
		},
		{
			name:         "Invalid slug",
			domain:       "domain",
			slug:         "nonexists",
			expectErr:    true,
			expectResult: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := m.LoadMetadata(tt.domain, tt.slug)

			if tt.expectErr {
				assert.Error(t, err, "error was expected")
			} else {
				assert.NoError(t, err, "error was not expected")
				assert.NotEmpty(t, result, "result was empty")
			}
		})
	}
}
