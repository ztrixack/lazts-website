package book

import (
	"lazts/internal/models"
	"lazts/internal/utils"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStats(t *testing.T) {
	content := `[
		{"title": "Book 1", "subtitle": "Subtitle 1", "description": "Description 1", "authors": ["Author 1"], "translator": "Translator 1", "publisher": "Publisher 1", "cover": "cover1.webp", "status": "completed", "catalog": "Catalog 1"},
		{"title": "Book 2", "subtitle": "Subtitle 2", "description": "Description 2", "authors": ["Author 2"], "translator": "Translator 2", "publisher": "Publisher 2", "cover": "cover2.webp", "status": "reading", "catalog": "Catalog 2"},
		{"title": "Book 3", "subtitle": "Subtitle 3", "description": "Description 3", "authors": ["Author 3"], "translator": "Translator 3", "publisher": "Publisher 3", "cover": "cover3.webp", "status": "unread", "catalog": "Catalog 3"},
		{"title": "Book 4", "subtitle": "Subtitle 4", "description": "Description 4", "authors": ["Author 4"], "translator": "Translator 4", "publisher": "Publisher 4", "cover": "cover4.webp", "status": "completed", "catalog": "Catalog 4"}
	]`

	tests := []struct {
		name          string
		contentDir    string
		setup         func(t *testing.T, dir string)
		expectedStats *models.BookStats
		expectedError bool
	}{
		{
			name:       "Successful stats calculation",
			contentDir: "./test_content",
			setup: func(t *testing.T, dir string) {
				utils.CreateTestFile(t, dir, "books/books.json", content)
			},
			expectedStats: &models.BookStats{
				Total:     4,
				Completed: 2,
				Reading:   1,
				Unread:    1,
				Pending:   2,
			},
			expectedError: false,
		},
		{
			name:       "List method returns error",
			contentDir: "./invalid",
			setup: func(t *testing.T, dir string) {
				// Do nothing, directory does not exist
			},
			expectedStats: nil,
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

			stats, err := s.GetStats()

			if tt.expectedError {
				assert.Error(t, err, "Expected an error")
			} else {
				assert.NoError(t, err, "Did not expect an error")
				assert.Equal(t, tt.expectedStats, stats, "Stats should match expected value")
			}
		})
	}
}
