package memo

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
			name:           "Successful size getting",
			expectedResult: 3,
			setup: func(t *testing.T, mock *markdown.Mock) {
				utils.CreateTestFile(t, CONTENT_DIR, "memos/slug-1/index.md", "")
				utils.CreateTestFile(t, CONTENT_DIR, "memos/slug-2/index.md", "")
				utils.CreateTestFile(t, CONTENT_DIR, "memos/slug-3/index.md", "")

				mock.On("LoadMetadata", "memos", "slug-1").Return(map[string]interface{}{}, nil).Once()
				mock.On("LoadMetadata", "memos", "slug-2").Return(map[string]interface{}{}, nil).Once()
				mock.On("LoadMetadata", "memos", "slug-3").Return(map[string]interface{}{}, nil).Once()
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
