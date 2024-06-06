package book

import (
	"lazts/internal/models"
	"lazts/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStats(t *testing.T) {
	const CONTENT_DIR = "test_get_stats"

	tests := []struct {
		name           string
		setup          func(t *testing.T)
		expectedError  bool
		expectedResult *models.BookStats
	}{
		{
			name:          "Successful stats calculation",
			expectedError: false,
			expectedResult: &models.BookStats{
				Total:     4,
				Completed: 2,
				Reading:   1,
				Unread:    1,
				Pending:   2,
			},
			setup: func(t *testing.T) {
				utils.CreateTestFile(t, CONTENT_DIR, "books/books.json", `[
					{"status": "completed"},
					{"status": "reading"},
					{"status": "unread"},
					{"status": "completed"}
				]`)
			},
		},
		{
			name:           "List method returns error",
			expectedError:  true,
			expectedResult: nil,
			setup: func(t *testing.T) {
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := setup(CONTENT_DIR)
			defer teardown(t, CONTENT_DIR)
			tt.setup(t)

			stats, err := s.GetStats()

			if tt.expectedError {
				assert.Error(t, err, "Expected an error")
			} else {
				assert.NoError(t, err, "Did not expect an error")
				assert.Equal(t, tt.expectedResult, stats, "Stats should match expected value")
			}
		})
	}
}
