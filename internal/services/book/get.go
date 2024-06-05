package book

import (
	"encoding/json"
	"lazts/internal/models"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func (s *service) Get() ([]models.Book, error) {
	if len(s.cache) != 0 {
		return s.cache, nil
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

			for i, book := range list {
				if book.Cover == "" {
					list[i].Cover = "https://picsum.photos/640/480"
					continue
				}
				name, _ := strings.CutSuffix(file.Name(), ".json")
				list[i].Cover, err = url.JoinPath("/static/contents/books", name, book.Cover)
				if err != nil {
					return nil, err
				}

				if book.Review == "" {
					list[i].Review = book.Description
				}
			}

			books = append(books, list...)
		}
	}

	s.cache = books
	return books, nil
}
