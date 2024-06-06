package book

import (
	"lazts/internal/utils"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSize(t *testing.T) {
	content := `[{},{},{},{},{}]`

	tests := []struct {
		name           string
		contentDir     string
		size           int
		setup          func(t *testing.T, dir string)
		teardown       func(t *testing.T, dir string)
		expectedResult int
	}{
		{
			name:       "Successful retrieval",
			contentDir: "test_data",
			size:       0,
			setup: func(t *testing.T, dir string) {
				os.Setenv("CONTENT_DIR", dir)
				utils.CreateTestFile(t, dir, "books/test.json", content)
			},
			teardown: func(t *testing.T, dir string) {
				os.Unsetenv("CONTENT_DIR")
				utils.RemoveTestDir(t, dir)
			},
			expectedResult: 5,
		},
		{
			name:       "Error from Get method",
			contentDir: "test_data",
			size:       0,
			setup: func(t *testing.T, dir string) {
				os.Setenv("CONTENT_DIR", dir)
			},
			teardown: func(t *testing.T, dir string) {
				os.Unsetenv("CONTENT_DIR")
				utils.RemoveTestDir(t, dir)
			},
			expectedResult: 0,
		},
		{
			name:       "Return cached size",
			contentDir: "test_data",
			size:       7,
			setup: func(t *testing.T, dir string) {
				os.Setenv("CONTENT_DIR", dir)
			},
			teardown: func(t *testing.T, dir string) {
				os.Unsetenv("CONTENT_DIR")
				utils.RemoveTestDir(t, dir)
			},
			expectedResult: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup(t, tt.contentDir)
			defer tt.teardown(t, tt.contentDir)

			r := New()
			r.size = tt.size
			result := r.GetSize()

			assert.Equal(t, tt.expectedResult, result, "Expected and actual results do not match")
		})
	}
}
