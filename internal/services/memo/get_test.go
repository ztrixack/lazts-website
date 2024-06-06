package memo

import (
	"lazts/internal/models"
	"lazts/internal/modules/markdown"
	"lazts/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	const CONTENT_DIR = "test_get"

	tests := []struct {
		name            string
		offset          uint
		limit           uint
		tag             string
		expectedError   bool
		expectedResults []models.Memo
		setup           func(t *testing.T, mock *markdown.Mock)
	}{
		{
			name:          "Successful retrieval with offset and limit",
			offset:        0,
			limit:         10,
			tag:           "",
			expectedError: false,
			expectedResults: []models.Memo{
				{
					Title:         "title",
					Excerpt:       "excerpt",
					FeaturedImage: "/static/contents/memos/slug-1/image.png",
					Link:          "/memos/tag1/slug-1",
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
			setup: func(t *testing.T, mock *markdown.Mock) {
				utils.CreateTestFile(t, CONTENT_DIR, "memos/slug-1/index.md", "")

				mock.On("LoadMetadata", "memos", "slug-1").Return(map[string]interface{}{
					"title":           "title",
					"slug":            "slug-1",
					"excerpt":         "excerpt",
					"featured_image":  "image.png",
					"published_at":    "2017-01-01",
					"published":       true,
					"last_updated_at": "2024-05-23",
					"tags":            []string{"tag1", "tag2"},
				}, nil).Once()
			},
		},
		{
			name:            "Directory does not exist",
			offset:          0,
			limit:           10,
			tag:             "",
			expectedError:   true,
			expectedResults: nil,
			setup: func(t *testing.T, mock *markdown.Mock) {
			},
		},
		{
			name:            "Metadata retrieval error",
			offset:          0,
			limit:           1,
			tag:             "",
			expectedError:   true,
			expectedResults: nil,
			setup: func(t *testing.T, mock *markdown.Mock) {
				utils.CreateTestFile(t, CONTENT_DIR, "memos/slug-1/index.md", "")

				mock.On("LoadMetadata", "memos", "slug-1").Return(map[string]interface{}{}, assert.AnError).Once()
			},
		},
		{
			name:          "Successful retrieval with offset, limit and tag",
			offset:        0,
			limit:         10,
			tag:           "tag2",
			expectedError: false,
			expectedResults: []models.Memo{
				{
					Title:         "title",
					Excerpt:       "excerpt",
					FeaturedImage: "/static/contents/memos/slug-1/image.png",
					Link:          "/memos/tag1/slug-1",
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
			setup: func(t *testing.T, mock *markdown.Mock) {
				utils.CreateTestFile(t, CONTENT_DIR, "memos/slug-1/index.md", "")
				utils.CreateTestFile(t, CONTENT_DIR, "memos/slug-2/index.md", "")

				mock.On("LoadMetadata", "memos", "slug-1").Return(map[string]interface{}{
					"title":           "title",
					"slug":            "slug-1",
					"excerpt":         "excerpt",
					"featured_image":  "image.png",
					"published_at":    "2017-01-01",
					"published":       true,
					"last_updated_at": "2024-05-23",
					"tags":            []string{"tag1", "tag2"},
				}, nil).Once()
				mock.On("LoadMetadata", "memos", "slug-2").Return(map[string]interface{}{"tag": []string{"tag1"}}, nil).Once()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, m := setup(CONTENT_DIR)
			defer teardown(t, CONTENT_DIR)
			tt.setup(t, m)

			memos, err := s.Get(tt.offset, tt.limit, tt.tag)

			if tt.expectedError {
				assert.Error(t, err, "Expected an error but did not get one")
			} else {
				assert.NoError(t, err, "Did not expect an error but got one")
				assert.ElementsMatch(t, tt.expectedResults, memos, "Expected and actual memos do not match")
			}

			m.AssertExpectations(t)
		})
	}
}
