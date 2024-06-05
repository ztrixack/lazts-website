package book

import (
	"lazts/internal/models"
	"lazts/internal/utils"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	content1 := `[{
		"title": "titleA",
		"subtitle": "subtitleA",
		"description": "descriptionA",
		"authors": ["author1", "author2"],
		"translator": "translatorA",
		"publisher": "publisherA",
		"cover": "cover.png",
		"status": "completed",
		"catalog": "catalog"
	},{
		"title": "titleB",
		"subtitle": "subtitleB",
		"description": "descriptionB",
		"authors": ["author3"],
		"translator": "translatorB",
		"publisher": "publisherB",
		"cover": "cover.webp",
		"status": "unread",
		"catalog": "catalog"
	}]`
	content2 := `[{
		"title": "titleC",
		"subtitle": "subtitleC",
		"description": "descriptionC",
		"authors": ["author4", "author5", "author6"],
		"translator": "translatorC",
		"publisher": "publisherC",
		"cover": "cover.jpg",
		"status": "reading",
		"catalog": "next"
	}]`

	tests := []struct {
		name          string
		contentDir    string
		setup         func(t *testing.T, dir string)
		teardown      func(t *testing.T, dir string)
		search        string
		catalog       string
		status        string
		expectedBooks []models.Book
		expectedError bool
	}{
		{
			name:       "Successful listing",
			contentDir: "test_content",
			setup: func(t *testing.T, dir string) {
				os.Setenv("CONTENT_DIR", dir)
				utils.CreateTestFile(t, dir, "books/book1.json", content1)
				utils.CreateTestFile(t, dir, "books/book2.json", content2)
			},
			teardown: func(t *testing.T, dir string) {
				os.Unsetenv("CONTENT_DIR")
				utils.RemoveTestDir(t, dir)
			},
			search:  "",
			catalog: "",
			status:  "",
			expectedBooks: []models.Book{
				{
					Title:       "titleA",
					Subtitle:    "subtitleA",
					Description: "descriptionA",
					Authors:     []string{"author1", "author2"},
					Translator:  "translatorA",
					Publisher:   "publisherA",
					Cover:       "/static/contents/books/book1/cover.png",
					Status:      "completed",
					Catalog:     "catalog",
					Review:      "descriptionA",
				},
				{
					Title:       "titleB",
					Subtitle:    "subtitleB",
					Description: "descriptionB",
					Authors:     []string{"author3"},
					Translator:  "translatorB",
					Publisher:   "publisherB",
					Cover:       "/static/contents/books/book1/cover.webp",
					Status:      "unread",
					Catalog:     "catalog",
					Review:      "descriptionB",
				},
				{
					Title:       "titleC",
					Subtitle:    "subtitleC",
					Description: "descriptionC",
					Authors:     []string{"author4", "author5", "author6"},
					Translator:  "translatorC",
					Publisher:   "publisherC",
					Cover:       "/static/contents/books/book2/cover.jpg",
					Status:      "reading",
					Catalog:     "next",
					Review:      "descriptionC",
				},
			},
			expectedError: false,
		},
		{
			name:       "Directory does not exist",
			contentDir: "invalid_content",
			setup: func(t *testing.T, dir string) {
				os.Setenv("CONTENT_DIR", dir)
			},
			teardown: func(t *testing.T, dir string) {
				os.Unsetenv("CONTENT_DIR")
				utils.RemoveTestDir(t, dir)
			},
			search:        "",
			catalog:       "",
			status:        "",
			expectedBooks: nil,
			expectedError: true,
		},
		{
			name:       "Filter by status",
			contentDir: "test_content",
			setup: func(t *testing.T, dir string) {
				os.Setenv("CONTENT_DIR", dir)
				utils.CreateTestFile(t, dir, "books/book1.json", content1)
			},
			teardown: func(t *testing.T, dir string) {
				os.Unsetenv("CONTENT_DIR")
				utils.RemoveTestDir(t, dir)
			},
			search:  "",
			catalog: "",
			status:  "completed",
			expectedBooks: []models.Book{
				{
					Title:       "titleA",
					Subtitle:    "subtitleA",
					Description: "descriptionA",
					Authors:     []string{"author1", "author2"},
					Translator:  "translatorA",
					Publisher:   "publisherA",
					Cover:       "/static/contents/books/book1/cover.png",
					Status:      "completed",
					Catalog:     "catalog",
					Review:      "descriptionA",
				},
			},
			expectedError: false,
		},
		{
			name:       "Filter by catalog",
			contentDir: "test_content",
			setup: func(t *testing.T, dir string) {
				os.Setenv("CONTENT_DIR", dir)
				utils.CreateTestFile(t, dir, "books/book1.json", content1)
				utils.CreateTestFile(t, dir, "books/book2.json", content2)
			},
			teardown: func(t *testing.T, dir string) {
				os.Unsetenv("CONTENT_DIR")
				utils.RemoveTestDir(t, dir)
			},
			search:  "",
			catalog: "catalog",
			status:  "",
			expectedBooks: []models.Book{
				{
					Title:       "titleA",
					Subtitle:    "subtitleA",
					Description: "descriptionA",
					Authors:     []string{"author1", "author2"},
					Translator:  "translatorA",
					Publisher:   "publisherA",
					Cover:       "/static/contents/books/book1/cover.png",
					Status:      "completed",
					Catalog:     "catalog",
					Review:      "descriptionA",
				},
				{
					Title:       "titleB",
					Subtitle:    "subtitleB",
					Description: "descriptionB",
					Authors:     []string{"author3"},
					Translator:  "translatorB",
					Publisher:   "publisherB",
					Cover:       "/static/contents/books/book1/cover.webp",
					Status:      "unread",
					Catalog:     "catalog",
					Review:      "descriptionB",
				},
			},
			expectedError: false,
		},
		{
			name:       "Filter by search term",
			contentDir: "test_content",
			setup: func(t *testing.T, dir string) {
				os.Setenv("CONTENT_DIR", dir)
				utils.CreateTestFile(t, dir, "books/book1.json", content1)
				utils.CreateTestFile(t, dir, "books/book2.json", content2)
			},
			teardown: func(t *testing.T, dir string) {
				os.Unsetenv("CONTENT_DIR")
				utils.RemoveTestDir(t, dir)
			},
			search:  "tlea",
			catalog: "",
			status:  "",
			expectedBooks: []models.Book{
				{
					Title:       "titleA",
					Subtitle:    "subtitleA",
					Description: "descriptionA",
					Authors:     []string{"author1", "author2"},
					Translator:  "translatorA",
					Publisher:   "publisherA",
					Cover:       "/static/contents/books/book1/cover.png",
					Status:      "completed",
					Catalog:     "catalog",
					Review:      "descriptionA",
				},
			},
			expectedError: false,
		},
		{
			name:       "Invalid JSON file",
			contentDir: "test_content",
			setup: func(t *testing.T, dir string) {
				os.Setenv("CONTENT_DIR", dir)
				utils.CreateTestFile(t, dir, "books/invalid.json", "{invalid json}")
			},
			teardown: func(t *testing.T, dir string) {
				os.Unsetenv("CONTENT_DIR")
				utils.RemoveTestDir(t, dir)
			},
			search:        "",
			catalog:       "",
			status:        "",
			expectedBooks: nil,
			expectedError: true,
		},
		{
			name:       "Empty JSON file",
			contentDir: "test_content",
			setup: func(t *testing.T, dir string) {
				os.Setenv("CONTENT_DIR", dir)
				utils.CreateTestFile(t, dir, "books/empty.json", "[]")
			},
			teardown: func(t *testing.T, dir string) {
				os.Unsetenv("CONTENT_DIR")
				utils.RemoveTestDir(t, dir)
			},
			search:        "",
			catalog:       "",
			status:        "",
			expectedBooks: []models.Book{},
			expectedError: false,
		},
		{
			name:       "Empty fields and cover defaults",
			contentDir: "test_content",
			setup: func(t *testing.T, dir string) {
				os.Setenv("CONTENT_DIR", dir)
				utils.CreateTestFile(t, dir, "books/book1.json", `[{
					"title": "title",
					"subtitle": "",
					"description": "",
					"authors": [],
					"translator": "",
					"publisher": "",
					"cover": "",
					"status": "completed",
					"catalog": "catalog"
				}]`)
			},
			teardown: func(t *testing.T, dir string) {
				os.Unsetenv("CONTENT_DIR")
				utils.RemoveTestDir(t, dir)
			},
			search:  "",
			catalog: "",
			status:  "",
			expectedBooks: []models.Book{
				{
					Title:       "title",
					Subtitle:    "",
					Description: "",
					Authors:     []string{},
					Translator:  "",
					Publisher:   "",
					Cover:       "https://picsum.photos/640/480",
					Status:      "completed",
					Catalog:     "catalog",
					Review:      "",
				},
			},
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup(t, tt.contentDir)
			defer tt.teardown(t, tt.contentDir)

			s := New()
			books, err := s.Get(tt.search, tt.catalog, tt.status)

			if tt.expectedError {
				assert.Error(t, err, "Expected an error but did not get one")
			} else {
				assert.NoError(t, err, "Did not expect an error but got one")
				assert.ElementsMatch(t, tt.expectedBooks, books, "Books should match expected value")
			}
		})
	}
}
