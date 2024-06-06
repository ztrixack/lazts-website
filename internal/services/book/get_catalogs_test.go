package book

import (
	"lazts/internal/models"
	"lazts/internal/utils"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCatalogs(t *testing.T) {
	content := `[{
		"catalog": "A"
	},{
		"catalog": "A"
	},{
		"catalog": "B"
	},{
		"catalog": "C"
	},{
		"catalog": "D"
	}]`

	tests := []struct {
		name           string
		contentDir     string
		cached         []models.Option
		setup          func(t *testing.T, dir string)
		teardown       func(t *testing.T, dir string)
		expectedResult []models.Option
		expectedError  bool
	}{
		{
			name:       "Successful retrieval",
			contentDir: "test_data",
			cached:     nil,
			setup: func(t *testing.T, dir string) {
				os.Setenv("CONTENT_DIR", dir)
				utils.CreateTestFile(t, dir, "books/test.json", content)
			},
			teardown: func(t *testing.T, dir string) {
				os.Unsetenv("CONTENT_DIR")
				utils.RemoveTestDir(t, dir)
			},
			expectedResult: []models.Option{
				{Key: "A", Value: "A"},
				{Key: "B", Value: "B"},
				{Key: "C", Value: "C"},
				{Key: "D", Value: "D"},
				{Key: "ทั้งหมด", Value: ""},
			},
			expectedError: false,
		},
		{
			name:       "Error from Get method",
			contentDir: "test_data",
			cached:     nil,
			setup: func(t *testing.T, dir string) {
				os.Setenv("CONTENT_DIR", dir)
			},
			teardown: func(t *testing.T, dir string) {
				os.Unsetenv("CONTENT_DIR")
				utils.RemoveTestDir(t, dir)
			},
			expectedResult: nil,
			expectedError:  true,
		},
		{
			name:       "Return cached catalogs",
			contentDir: "test_data",
			cached: []models.Option{
				{Key: "ทั้งหมด", Value: ""},
				{Key: "CachedCatalog", Value: "CachedCatalog"},
			},
			setup: func(t *testing.T, dir string) {
				os.Setenv("CONTENT_DIR", dir)
			},
			teardown: func(t *testing.T, dir string) {
				os.Unsetenv("CONTENT_DIR")
				utils.RemoveTestDir(t, dir)
			},
			expectedResult: []models.Option{
				{Key: "CachedCatalog", Value: "CachedCatalog"},
				{Key: "ทั้งหมด", Value: ""},
			},
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup(t, tt.contentDir)
			defer tt.teardown(t, tt.contentDir)

			r := New()
			r.catalogs = tt.cached
			result, err := r.GetCatalogs()

			if tt.expectedError {
				assert.Error(t, err, "Expected an error but did not get one")
			} else {
				assert.NoError(t, err, "Did not expect an error but got one")
				assert.ElementsMatch(t, tt.expectedResult, result, "Expected and actual results do not match")
			}
		})
	}
}
