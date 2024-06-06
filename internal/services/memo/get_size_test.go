package memo

import (
	"lazts/internal/modules/markdown"
	"lazts/internal/utils"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSize(t *testing.T) {
	tests := []struct {
		name           string
		contentDir     string
		size           int
		setup          func(t *testing.T, dir string, mock *markdown.Mock)
		teardown       func(t *testing.T, dir string)
		expectedResult int
	}{
		{
			name:       "Successful retrieval",
			contentDir: "test_content",
			size:       0,
			setup: func(t *testing.T, dir string, mock *markdown.Mock) {
				os.Setenv("CONTENT_DIR", dir)
				utils.CreateTestFile(t, dir, "memos/00000000-slug-1/index.md", "some content")
				utils.CreateTestFile(t, dir, "memos/00000000-slug-2/index.md", "some content")
				utils.CreateTestFile(t, dir, "memos/00000000-slug-3/index.md", "some content")

				mock.On("ReadFile", "00000000-slug-1", dir+"/memos/00000000-slug-1/index.md").Return([]byte{}, nil).Once()
				mock.On("ToMetadata", "00000000-slug-1", []byte{}).Return(map[string]interface{}{}, nil).Once()
				mock.On("ReadFile", "00000000-slug-2", dir+"/memos/00000000-slug-2/index.md").Return([]byte{}, nil).Once()
				mock.On("ToMetadata", "00000000-slug-2", []byte{}).Return(map[string]interface{}{}, nil).Once()
				mock.On("ReadFile", "00000000-slug-3", dir+"/memos/00000000-slug-3/index.md").Return([]byte{}, nil).Once()
				mock.On("ToMetadata", "00000000-slug-3", []byte{}).Return(map[string]interface{}{}, nil).Once()
			},
			teardown: func(t *testing.T, dir string) {
				os.Unsetenv("CONTENT_DIR")
				utils.RemoveTestDir(t, dir)
			},
			expectedResult: 3,
		},
		{
			name:       "Error from Get method",
			contentDir: "test_data",
			size:       0,
			setup: func(t *testing.T, dir string, mock *markdown.Mock) {
				os.Setenv("CONTENT_DIR", dir)
			},
			teardown: func(t *testing.T, dir string) {
				os.Unsetenv("CONTENT_DIR")
				utils.RemoveTestDir(t, dir)
			},
			expectedResult: 0,
		},
		{
			name:       "Return cached size",
			contentDir: "test_data",
			size:       7,
			setup: func(t *testing.T, dir string, mock *markdown.Mock) {
				os.Setenv("CONTENT_DIR", dir)
			},
			teardown: func(t *testing.T, dir string) {
				os.Unsetenv("CONTENT_DIR")
				utils.RemoveTestDir(t, dir)
			},
			expectedResult: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			markdownMock := new(markdown.Mock)
			tt.setup(t, tt.contentDir, markdownMock)
			defer tt.teardown(t, tt.contentDir)

			r := New(markdownMock)
			r.cache["SIZE"] = tt.size
			result := r.GetSize()

			assert.Equal(t, tt.expectedResult, result, "Expected and actual results do not match")
			markdownMock.AssertExpectations(t)
		})
	}
}
