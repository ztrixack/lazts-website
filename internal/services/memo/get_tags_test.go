package memo

import (
	"lazts/internal/models"
	"lazts/internal/modules/markdown"
	"lazts/internal/utils"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTags(t *testing.T) {
	tests := []struct {
		name           string
		contentDir     string
		cached         []models.Tag
		setup          func(t *testing.T, dir string, mock *markdown.Mock)
		teardown       func(t *testing.T, dir string)
		expectedResult []models.Tag
		expectedError  bool
	}{
		{
			name:       "Successful retrieval",
			contentDir: "test_content",
			cached:     nil,
			setup: func(t *testing.T, dir string, mock *markdown.Mock) {
				os.Setenv("CONTENT_DIR", dir)
				utils.CreateTestFile(t, dir, "memos/00000000-slug-1/index.md", "some content")
				utils.CreateTestFile(t, dir, "memos/00000000-slug-2/index.md", "some content")
				utils.CreateTestFile(t, dir, "memos/00000000-slug-3/index.md", "some content")

				mock.On("LoadMetadata", "memos", "00000000-slug-1").Return(map[string]interface{}{"tags": []string{"A"}}, nil).Once()
				mock.On("LoadMetadata", "memos", "00000000-slug-2").Return(map[string]interface{}{"tags": []string{"A", "B"}}, nil).Once()
				mock.On("LoadMetadata", "memos", "00000000-slug-3").Return(map[string]interface{}{"tags": []string{"C", "D"}}, nil).Once()
			},
			teardown: func(t *testing.T, dir string) {
				os.Unsetenv("CONTENT_DIR")
				utils.RemoveTestDir(t, dir)
			},
			expectedResult: []models.Tag{
				{Name: "A", Link: "/memos/a", Count: 2},
				{Name: "B", Link: "/memos/b", Count: 1},
				{Name: "C", Link: "/memos/c", Count: 1},
				{Name: "D", Link: "/memos/d", Count: 1},
			},
			expectedError: false,
		},
		{
			name:       "Error from Get method",
			contentDir: "test_data",
			cached:     nil,
			setup: func(t *testing.T, dir string, mock *markdown.Mock) {
				os.Setenv("CONTENT_DIR", dir)
			},
			teardown: func(t *testing.T, dir string) {
				os.Unsetenv("CONTENT_DIR")
				utils.RemoveTestDir(t, dir)
			},
			expectedResult: nil,
			expectedError:  true,
		},
		{
			name:       "Return cached tags",
			contentDir: "test_data",
			cached: []models.Tag{
				{Name: "A", Count: 2},
				{Name: "B", Count: 1},
				{Name: "C", Count: 1},
				{Name: "D", Count: 1},
			},
			setup: func(t *testing.T, dir string, mock *markdown.Mock) {
				os.Setenv("CONTENT_DIR", dir)
			},
			teardown: func(t *testing.T, dir string) {
				os.Unsetenv("CONTENT_DIR")
				utils.RemoveTestDir(t, dir)
			},
			expectedResult: []models.Tag{
				{Name: "A", Count: 2},
				{Name: "B", Count: 1},
				{Name: "C", Count: 1},
				{Name: "D", Count: 1},
			},
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			markdownMock := new(markdown.Mock)
			tt.setup(t, tt.contentDir, markdownMock)
			defer tt.teardown(t, tt.contentDir)

			r := New(markdownMock)
			r.cache["TAGS"] = tt.cached
			result, err := r.GetTags()

			if tt.expectedError {
				assert.Error(t, err, "Expected an error but did not get one")
			} else {
				assert.NoError(t, err, "Did not expect an error but got one")
				assert.ElementsMatch(t, tt.expectedResult, result, "Expected and actual memos do not match")
			}

			markdownMock.AssertExpectations(t)
		})
	}
}
