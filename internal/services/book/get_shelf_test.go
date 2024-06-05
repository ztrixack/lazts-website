package book

import (
	"lazts/internal/utils"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetShelf(t *testing.T) {
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
		rows          []int
		expectedShelf int
		expectedError bool
	}{
		{
			name:       "Successful book listing",
			contentDir: "test_content",
			setup: func(t *testing.T, dir string) {
				utils.CreateTestFile(t, dir, "books/book1.json", content)
				utils.CreateTestFile(t, dir, "books/book2.json", content)
			},
			rows:          []int{2, 3, 4},
			expectedShelf: 3,
			expectedError: false,
		},
		{
			name:       "Empty book listing",
			contentDir: "test_content",
			setup: func(t *testing.T, dir string) {
				utils.CreateTestFile(t, dir, "books/book1.json", content)
				utils.CreateTestFile(t, dir, "books/book2.json", content)
			},
			rows:          []int{},
			expectedShelf: 0,
			expectedError: false,
		},
		{
			name:       "Nil book listing",
			contentDir: "test_content",
			setup: func(t *testing.T, dir string) {
				utils.CreateTestFile(t, dir, "books/book1.json", content)
				utils.CreateTestFile(t, dir, "books/book2.json", content)
			},
			rows:          nil,
			expectedShelf: 0,
			expectedError: false,
		},
		{
			name:       "Error book listing",
			contentDir: "invalid_content",
			setup: func(t *testing.T, dir string) {
				// Do nothing, directory does not exist
			},
			rows:          nil,
			expectedShelf: 0,
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

			shelf, err := s.GetShelf(tt.rows...)

			if tt.expectedError {
				assert.Error(t, err, "Expected an error")
			} else {
				assert.NoError(t, err, "Did not expect an error")
				assert.Equal(t, tt.expectedShelf, len(shelf), "The number of shelves should match the expected value")

				for i, row := range shelf {
					assert.Equal(t, tt.rows[i], len(row), "The number of books in the row should match the expected value")
				}
			}
		})
	}
}
