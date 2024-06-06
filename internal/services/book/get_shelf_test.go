package book

import (
	"lazts/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetShelf(t *testing.T) {
	const CONTENT_DIR = "test_get_shelf"

	tests := []struct {
		name           string
		rows           []int
		expectedError  bool
		expectedResult int
		setup          func(t *testing.T)
	}{
		{
			name:           "Successful book listing",
			rows:           []int{2, 3, 4},
			expectedError:  false,
			expectedResult: 3,
			setup: func(t *testing.T) {
				utils.CreateTestFile(t, CONTENT_DIR, "books/book1.json", "[{},{},{},{},{}]")
				utils.CreateTestFile(t, CONTENT_DIR, "books/book2.json", "[{},{},{},{},{}]")
			},
		},
		{
			name:           "Empty book listing",
			rows:           []int{},
			expectedError:  false,
			expectedResult: 0,
			setup: func(t *testing.T) {
				utils.CreateTestFile(t, CONTENT_DIR, "books/book1.json", "[]")
				utils.CreateTestFile(t, CONTENT_DIR, "books/book2.json", "[]")
			},
		},
		{
			name:          "Invalid book listing",
			rows:          nil,
			expectedError: true,
			setup: func(t *testing.T) {
				utils.CreateTestFile(t, CONTENT_DIR, "books/book1.json", "invalid json")
			},
		},
		{
			name:          "Empty book listing",
			rows:          nil,
			expectedError: true,
			setup: func(t *testing.T) {
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := setup(CONTENT_DIR)
			defer teardown(t, CONTENT_DIR)
			tt.setup(t)

			shelf, err := s.GetShelf(tt.rows...)

			if tt.expectedError {
				assert.Error(t, err, "Expected an error")
			} else {
				assert.NoError(t, err, "Did not expect an error")
				assert.Equal(t, tt.expectedResult, len(shelf), "The number of shelves should match the expected value")

				for i, row := range shelf {
					assert.Equal(t, tt.rows[i], len(row), "The number of books in the row should match the expected value")
				}
			}
		})
	}
}
