package vacation

import (
	"lazts/internal/modules/markdown"
	"lazts/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSize(t *testing.T) {
	const CONTENT_DIR = "test_get_size"

	tests := []struct {
		name           string
		setup          func(t *testing.T, mock *markdown.Mock)
		expectedResult int
	}{
		{
			name:           "Successful retrieval",
			expectedResult: 3,
			setup: func(t *testing.T, mock *markdown.Mock) {
				utils.CreateTestFile(t, CONTENT_DIR, "vacations/slug-1/index.md", "")
				utils.CreateTestFile(t, CONTENT_DIR, "vacations/slug-2/index.md", "")
				utils.CreateTestFile(t, CONTENT_DIR, "vacations/slug-3/index.md", "")

				mock.On("LoadMetadata", "vacations", "slug-1").Return(map[string]interface{}{"published": true}, nil).Once()
				mock.On("LoadMetadata", "vacations", "slug-2").Return(map[string]interface{}{"published": true}, nil).Once()
				mock.On("LoadMetadata", "vacations", "slug-3").Return(map[string]interface{}{"published": true}, nil).Once()
			},
		},
		{
			name:           "Error from Get method",
			expectedResult: 0,
			setup: func(t *testing.T, mock *markdown.Mock) {
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, m := setup(CONTENT_DIR)
			defer teardown(t, CONTENT_DIR)
			tt.setup(t, m)

			result := s.GetSize()

			assert.Equal(t, tt.expectedResult, result, "Expected and actual results do not match")
			m.AssertExpectations(t)
		})
	}
}
