package vacation

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
		name           string
		location       string
		expectedError  bool
		expectedResult []models.Vacation
		setup          func(t *testing.T, mock *markdown.Mock)
	}{
		{
			name:          "Successful listing",
			location:      "",
			expectedError: false,
			expectedResult: []models.Vacation{
				{
					Title:            "title",
					Excerpt:          "excerpt",
					Location:         "location",
					DateTimeISO:      "2017-01-01T00:00:00Z",
					DateTimeReadable: "วันอาทิตย์ที่ 1 มกราคม 2017",
					FeaturedImage:    "/static/contents/vacations/slug-1/image.png",
					Link:             "/vacations/slug-1",
				},
			},
			setup: func(t *testing.T, mock *markdown.Mock) {
				utils.CreateTestFile(t, CONTENT_DIR, "vacations/slug-1/index.md", "")

				mock.On("LoadMetadata", "vacations", "slug-1").Return(map[string]interface{}{
					"title":           "title",
					"slug":            "slug-1",
					"excerpt":         "excerpt",
					"location":        "location",
					"date_from":       "2017-01-01",
					"date_to":         "2017-01-01",
					"featured_image":  "image.png",
					"published_at":    "2017-01-01",
					"published":       true,
					"last_updated_at": "2024-05-23",
				}, nil).Once()
			},
		},
		{
			name:          "Successful with location",
			location:      "location",
			expectedError: false,
			expectedResult: []models.Vacation{
				{
					Title:            "title",
					Excerpt:          "excerpt",
					Location:         "location",
					DateTimeISO:      "2017-01-01T00:00:00Z",
					DateTimeReadable: "วันอาทิตย์ที่ 1 มกราคม 2017",
					FeaturedImage:    "/static/contents/vacations/slug-1/image.png",
					Link:             "/vacations/slug-1",
				},
			},
			setup: func(t *testing.T, mock *markdown.Mock) {
				utils.CreateTestFile(t, CONTENT_DIR, "vacations/slug-1/index.md", "")
				utils.CreateTestFile(t, CONTENT_DIR, "vacations/slug-2/index.md", "")

				mock.On("LoadMetadata", "vacations", "slug-1").Return(map[string]interface{}{
					"title":           "title",
					"slug":            "slug-1",
					"excerpt":         "excerpt",
					"location":        "location",
					"date_from":       "2017-01-01",
					"date_to":         "2017-01-01",
					"featured_image":  "image.png",
					"published_at":    "2017-01-01",
					"published":       true,
					"last_updated_at": "2024-05-23",
				}, nil).Once()
				mock.On("LoadMetadata", "vacations", "slug-2").Return(map[string]interface{}{
					"location": "not-this-location",
				}, nil).Once()
			},
		},
		{
			name:           "Directory does not exist",
			location:       "",
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

			books, err := s.Get(tt.location)

			if tt.expectedError {
				assert.Error(t, err, "Expected an error")
			} else {
				assert.NoError(t, err, "Did not expect an error")
				assert.ElementsMatch(t, tt.expectedResult, books, "Books should match expected value")
			}
		})
	}
}
