package markdown

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarkdown_LoadMetadata(t *testing.T) {
	m := setup(t)
	defer teardown(t)

	tests := []struct {
		name           string
		domain         string
		slug           string
		expectedErr    bool
		expectedResult map[string]interface{}
	}{
		{
			name:           "Successful metadata loading",
			domain:         "domain",
			slug:           "slug",
			expectedErr:    false,
			expectedResult: map[string]interface{}{"meta": "test", "read_time": 0},
		},
		{
			name:           "Invalid domain",
			domain:         "nonexists",
			slug:           "slug",
			expectedErr:    true,
			expectedResult: nil,
		},
		{
			name:           "Invalid slug",
			domain:         "domain",
			slug:           "nonexists",
			expectedErr:    true,
			expectedResult: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := m.LoadMetadata(tt.domain, tt.slug)

			if tt.expectedErr {
				assert.Error(t, err, "error was expected")
			} else {
				assert.NoError(t, err, "error was not expected")
				assert.NotEmpty(t, result, "result was empty")
			}
		})
	}
}
