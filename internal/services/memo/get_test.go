package memo

import (
	"lazts/internal/models"
	"lazts/internal/modules/markdown"
	"lazts/internal/utils"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	content := map[string]interface{}{
		"title":           "title",
		"slug":            "00000000-slug-1",
		"excerpt":         "excerpt",
		"featured_image":  "image.png",
		"published_at":    "2017-01-01",
		"published":       true,
		"last_updated_at": "2024-05-23",
		"tags":            []string{"tag1", "tag2"},
	}

	tests := []struct {
		name          string
		contentDir    string
		setup         func(t *testing.T, dir string, mock *markdown.Mock)
		teardown      func(t *testing.T, dir string)
		offset        uint
		limit         uint
		tag           string
		expectedMemos []models.Memo
		expectedError bool
	}{
		{
			name:       "Successful retrieval with offset and limit",
			contentDir: "test_content",
			setup: func(t *testing.T, dir string, mock *markdown.Mock) {
				os.Setenv("CONTENT_DIR", dir)
				utils.CreateTestFile(t, dir, "memos/00000000-slug-1/index.md", "some content")

				mock.On("LoadMetadata", "memos", "00000000-slug-1").Return(content, nil).Once()
			},
			teardown: func(t *testing.T, dir string) {
				os.Unsetenv("CONTENT_DIR")
				utils.RemoveTestDir(t, dir)
			},
			offset: 0,
			limit:  10,
			tag:    "",
			expectedMemos: []models.Memo{
				{
					Title:         "title",
					Excerpt:       "excerpt",
					FeaturedImage: "/static/contents/memos/00000000-slug-1/image.png",
					Link:          "/memos/tag1/00000000-slug-1",
					Tags: []models.Tag{
						{Name: "tag1", Link: "/memos/tag1", Count: 1},
						{Name: "tag2", Link: "/memos/tag2", Count: 1},
					},
					DateTimeISO:         "2017-01-01T00:00:00Z",
					DateTimeReadable:    "01 มกราคม 2017",
					LastUpdatedISO:      "2024-05-23T00:00:00Z",
					LastUpdatedReadable: "23 พฤษภาคม 2024",
					DayMonth:            "01 Jan",
					Year:                "2017",
					ReadTime:            0,
				},
			},
			expectedError: false,
		},
		{
			name:       "Directory does not exist",
			contentDir: "invalid_content",
			setup: func(t *testing.T, dir string, mock *markdown.Mock) {
				os.Setenv("CONTENT_DIR", dir)
			},
			teardown: func(t *testing.T, dir string) {
				os.Unsetenv("CONTENT_DIR")
				utils.RemoveTestDir(t, dir)
			},
			offset:        0,
			limit:         10,
			tag:           "",
			expectedMemos: nil,
			expectedError: true,
		},
		{
			name:       "Metadata retrieval error",
			contentDir: "test_content",
			setup: func(t *testing.T, dir string, mock *markdown.Mock) {
				os.Setenv("CONTENT_DIR", dir)
				utils.CreateTestFile(t, dir, "memos/00000000-slug-1/index.md", "some content")

				mock.On("LoadMetadata", "memos", "00000000-slug-1").Return(map[string]interface{}{}, assert.AnError).Once()
			},
			teardown: func(t *testing.T, dir string) {
				os.Unsetenv("CONTENT_DIR")
				utils.RemoveTestDir(t, dir)
			},
			offset:        0,
			limit:         1,
			tag:           "",
			expectedMemos: nil,
			expectedError: true,
		},
		{
			name:       "Successful retrieval with offset, limit and tag",
			contentDir: "test_content",
			setup: func(t *testing.T, dir string, mock *markdown.Mock) {
				os.Setenv("CONTENT_DIR", dir)
				utils.CreateTestFile(t, dir, "memos/00000000-slug-1/index.md", "some content")
				utils.CreateTestFile(t, dir, "memos/00000000-slug-2/index.md", "some content")

				mock.On("LoadMetadata", "memos", "00000000-slug-1").Return(content, nil).Once()
				mock.On("LoadMetadata", "memos", "00000000-slug-2").Return(map[string]interface{}{"tag": []string{"tag1"}}, nil).Once()
			},
			teardown: func(t *testing.T, dir string) {
				os.Unsetenv("CONTENT_DIR")
				utils.RemoveTestDir(t, dir)
			},
			offset: 0,
			limit:  10,
			tag:    "tag2",
			expectedMemos: []models.Memo{
				{
					Title:         "title",
					Excerpt:       "excerpt",
					FeaturedImage: "/static/contents/memos/00000000-slug-1/image.png",
					Link:          "/memos/tag1/00000000-slug-1",
					Tags: []models.Tag{
						{Name: "tag1", Link: "/memos/tag1", Count: 1},
						{Name: "tag2", Link: "/memos/tag2", Count: 1},
					},
					DateTimeISO:         "2017-01-01T00:00:00Z",
					DateTimeReadable:    "01 มกราคม 2017",
					LastUpdatedISO:      "2024-05-23T00:00:00Z",
					LastUpdatedReadable: "23 พฤษภาคม 2024",
					DayMonth:            "01 Jan",
					Year:                "2017",
					ReadTime:            0,
				},
			},
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			markdownMock := new(markdown.Mock)
			tt.setup(t, tt.contentDir, markdownMock)
			defer tt.teardown(t, tt.contentDir)

			s := New(markdownMock)
			memos, err := s.Get(tt.offset, tt.limit, tt.tag)

			if tt.expectedError {
				assert.Error(t, err, "Expected an error but did not get one")
			} else {
				assert.NoError(t, err, "Did not expect an error but got one")
				assert.ElementsMatch(t, tt.expectedMemos, memos, "Expected and actual memos do not match")
			}

			markdownMock.AssertExpectations(t)
		})
	}
}
