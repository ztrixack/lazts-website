package memo

import (
	"lazts/internal/models"
	"lazts/internal/modules/markdown"
	"lazts/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTags(t *testing.T) {
	const CONTENT_DIR = "test_get_tags"

	tests := []struct {
		name           string
		expectedError  bool
		expectedResult []models.Tag
		setup          func(t *testing.T, mock *markdown.Mock)
	}{
		{
			name:          "Successful retrieval",
			expectedError: false,
			expectedResult: []models.Tag{
				{Name: "A", Link: "/memos/a", Count: 4},
				{Name: "B", Link: "/memos/b", Count: 1},
				{Name: "C", Link: "/memos/c", Count: 2},
				{Name: "D", Link: "/memos/d", Count: 1},
			},
			setup: func(t *testing.T, mock *markdown.Mock) {
				utils.CreateTestFile(t, CONTENT_DIR, "memos/slug-1/index.md", "")
				utils.CreateTestFile(t, CONTENT_DIR, "memos/slug-2/index.md", "")
				utils.CreateTestFile(t, CONTENT_DIR, "memos/slug-3/index.md", "")
				utils.CreateTestFile(t, CONTENT_DIR, "memos/slug-4/index.md", "")

				mock.On("LoadMetadata", "memos", "slug-1").Return(map[string]interface{}{"published": true, "tags": []string{"A"}}, nil).Once()
				mock.On("LoadMetadata", "memos", "slug-2").Return(map[string]interface{}{"published": true, "tags": []string{"A", "B"}}, nil).Once()
				mock.On("LoadMetadata", "memos", "slug-3").Return(map[string]interface{}{"published": true, "tags": []string{"A", "C"}}, nil).Once()
				mock.On("LoadMetadata", "memos", "slug-4").Return(map[string]interface{}{"published": true, "tags": []string{"A", "C", "D"}}, nil).Once()
			},
		},
		{
			name:           "Error from Get method",
			expectedError:  true,
			expectedResult: nil,
			setup: func(t *testing.T, mock *markdown.Mock) {
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, m := setup(CONTENT_DIR)
			defer teardown(t, CONTENT_DIR)
			tt.setup(t, m)

			result, err := s.GetTags()

			if tt.expectedError {
				assert.Error(t, err, "Expected an error but did not get one")
			} else {
				assert.NoError(t, err, "Did not expect an error but got one")
				assert.ElementsMatch(t, tt.expectedResult, result, "Expected and actual memos do not match")
			}

			m.AssertExpectations(t)
		})
	}
}
