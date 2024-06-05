package book

import (
	"encoding/json"
	"lazts/internal/models"
	"os"
	"path/filepath"
	"strings"
)

func (s *service) Get(search, catalog, status string) ([]models.Book, error) {
	cpath := search + catalog + status
	if len(s.caches[cpath]) != 0 {
		return s.caches[cpath], nil
	}

	files, err := os.ReadDir(filepath.Join(s.config.ContentDir, "books"))
	if err != nil {
		return nil, err
	}

	books := make([]models.Book, 0)
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") {
			bytes, err := os.ReadFile(filepath.Join(s.config.ContentDir, "books", file.Name()))
			if err != nil {
				return nil, err
			}

			var list []models.Book
			if err := json.Unmarshal(bytes, &list); err != nil {
				return nil, err
			}

			for _, book := range list {
				if status != "" && book.Status != status {
					continue
				}

				if catalog != "" && book.Catalog != catalog {
					continue
				}

				if search != "" && !strings.Contains(strings.ToLower(book.Title), strings.ToLower(search)) {
					continue
				}

				if book.Cover == "" {
					book.Cover = "https://picsum.photos/640/480"
				} else {
					name := strings.TrimSuffix(file.Name(), ".json")
					book.Cover = "/static/contents/books/" + name + "/" + book.Cover
				}

				if book.Review == "" {
					book.Review = book.Description
				}

				books = append(books, book)
			}
		}
	}

	s.caches[cpath] = books
	return books, nil
}
