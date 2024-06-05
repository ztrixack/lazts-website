package book

import (
	"lazts/internal/models"
	"lazts/internal/utils"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	content := `[{
		"title": "title",
		"subtitle": "subtitle",
		"description": "description",
		"authors": ["author1", "author2"],
		"translator": "translator",
		"publisher": "publisher",
		"cover": "cover.webp",
		"status": "completed",
		"catalog": "catalog"
	}]`

	tests := []struct {
		name          string
		contentDir    string
		setup         func(t *testing.T, dir string)
		expectedBooks []models.Book
		expectedError bool
	}{
		{
			name:       "Successful listing",
			contentDir: "test_content",
			setup: func(t *testing.T, dir string) {
				utils.CreateTestFile(t, dir, "books/book1.json", content)
				utils.CreateTestFile(t, dir, "books/book2.json", content)
			},
			expectedBooks: []models.Book{
				{
					Number:      "",
					Title:       "title",
					Subtitle:    "subtitle",
					Description: "description",
					Authors:     []string{"author1", "author2"},
					Translator:  "translator",
					Publisher:   "publisher",
					Cover:       "/static/contents/books/book1/cover.webp",
					Status:      "completed",
					Catalog:     "catalog",
					Review:      "description",
				},
				{
					Number:      "",
					Title:       "title",
					Subtitle:    "subtitle",
					Description: "description",
					Authors:     []string{"author1", "author2"},
					Translator:  "translator",
					Publisher:   "publisher",
					Cover:       "/static/contents/books/book2/cover.webp",
					Status:      "completed",
					Catalog:     "catalog",
					Review:      "description",
				},
			},
			expectedError: false,
		},
		{
			name:       "Directory does not exist",
			contentDir: "invalid_content",
			setup: func(t *testing.T, dir string) {
				// Do nothing, directory does not exist
			},
			expectedBooks: nil,
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup(t, tt.contentDir)
			}
			os.Setenv("CONTENT_DIR", tt.contentDir)
			defer os.Unsetenv("CONTENT_DIR")
			defer utils.RemoveTestDir(t, tt.contentDir)

			s := New()

			books, err := s.List()

			if tt.expectedError {
				assert.Error(t, err, "Expected an error")
			} else {
				assert.NoError(t, err, "Did not expect an error")
				assert.ElementsMatch(t, tt.expectedBooks, books, "Books should match expected value")
			}
		})
	}
}
