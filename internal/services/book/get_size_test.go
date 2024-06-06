package book

import (
	"lazts/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSize(t *testing.T) {
	const CONTENT_DIR = "test_get_size"

	tests := []struct {
		name           string
		setup          func(t *testing.T)
		expectedResult int
	}{
		{
			name:           "Successful retrieval",
			expectedResult: 5,
			setup: func(t *testing.T) {
				utils.CreateTestFile(t, CONTENT_DIR, "books/book1.json", "[{},{},{},{},{}]")
			},
		},
		{
			name:           "Error from Get method",
			expectedResult: 0,
			setup: func(t *testing.T) {
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := setup(CONTENT_DIR)
			defer teardown(t, CONTENT_DIR)
			tt.setup(t)

			result := s.GetSize()

			assert.Equal(t, tt.expectedResult, result, "Expected and actual results do not match")
		})
	}
}
