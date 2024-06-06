package book

import (
	"lazts/internal/models"
	"lazts/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCatalogs(t *testing.T) {
	const CONTENT_DIR = "test_get_catalogs"

	tests := []struct {
		name           string
		setup          func(t *testing.T)
		expectedError  bool
		expectedResult []models.Option
	}{
		{
			name:          "Successful retrieval",
			expectedError: false,
			expectedResult: []models.Option{
				{Key: "A", Value: "A"},
				{Key: "B", Value: "B"},
				{Key: "C", Value: "C"},
				{Key: "D", Value: "D"},
				{Key: "ทั้งหมด", Value: ""},
			},
			setup: func(t *testing.T) {
				utils.CreateTestFile(t, CONTENT_DIR, "books/test.json", `[
					{"catalog": "A"},
					{"catalog": "A"},
					{"catalog": "B"},
					{"catalog": "C"},
					{"catalog": "D"}]`)
			},
		},
		{
			name:           "Error from Get method",
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

			result, err := s.GetCatalogs()

			if tt.expectedError {
				assert.Error(t, err, "Expected an error but did not get one")
			} else {
				assert.NoError(t, err, "Did not expect an error but got one")
				assert.ElementsMatch(t, tt.expectedResult, result, "Expected and actual results do not match")
			}
		})
	}
}
