package book

import (
	"lazts/internal/models"
)

func (s *service) GetCatalogs() ([]models.Option, error) {
	if s.catalogs != nil {
		return s.catalogs, nil
	}

	books, err := s.Get("", "", "")
	if err != nil {
		return nil, err
	}

	catalogs := models.Options{models.Option{Key: "ทั้งหมด", Value: ""}}
	for _, book := range books {
		catalogs = catalogs.AppendUnique(book.Catalog)
	}

	s.size = len(books)
	s.catalogs = catalogs.Sort()
	return s.catalogs, nil
}
