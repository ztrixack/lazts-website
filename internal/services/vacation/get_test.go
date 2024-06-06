package vacation

import (
	"lazts/internal/models"
	"lazts/internal/modules/markdown"
	"lazts/internal/utils"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	tests := []struct {
		name              string
		contentDir        string
		setup             func(t *testing.T, dir string, mock *markdown.Mock)
		expectedVacations []models.Vacation
		expectedError     bool
	}{
		{
			name:       "Successful listing",
			contentDir: "test_content",
			setup: func(t *testing.T, dir string, mock *markdown.Mock) {
				utils.CreateTestFile(t, dir, "vacations/0000-slug/index.md", "some content")

				mock.On("LoadMetadata", "vacations", "0000-slug").Return(map[string]interface{}{
					"title":           "title",
					"slug":            "0000-slug",
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
			expectedVacations: []models.Vacation{
				{
					Title:            "title",
					Excerpt:          "excerpt",
					Location:         "location",
					DateTimeISO:      "2017-01-01T00:00:00Z",
					DateTimeReadable: "วันอาทิตย์ที่ 1 มกราคม 2017",
					FeaturedImage:    "/static/contents/vacations/0000-slug/image.png",
					Link:             "/vacations/0000-slug",
				},
			},
			expectedError: false,
		},
		{
			name:       "Directory does not exist",
			contentDir: "invalid_content",
			setup: func(t *testing.T, dir string, mock *markdown.Mock) {
				// Do nothing, directory does not exist
			},
			expectedVacations: nil,
			expectedError:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			markdownMock := new(markdown.Mock)
			if tt.setup != nil {
				tt.setup(t, tt.contentDir, markdownMock)
			}
			os.Setenv("CONTENT_DIR", tt.contentDir)
			defer os.Unsetenv("CONTENT_DIR")
			defer utils.RemoveTestDir(t, tt.contentDir)

			s := New(markdownMock)

			books, err := s.Get("location")

			if tt.expectedError {
				assert.Error(t, err, "Expected an error")
			} else {
				assert.NoError(t, err, "Did not expect an error")
				assert.ElementsMatch(t, tt.expectedVacations, books, "Books should match expected value")
			}
		})
	}
}
