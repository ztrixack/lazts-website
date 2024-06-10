package book

import (
	"lazts/internal/models"
)

func (s *service) GetCatalogs() ([]models.Option, error) {
	const KEY = "CATALOG"
	if value, found := s.cache.Get(KEY); found {
		return value.([]models.Option), nil
	}

	books, err := s.Get("", "", "")
	if err != nil {
		return nil, err
	}

	catalogs := models.Options{models.Option{Key: "ทั้งหมด", Value: ""}}
	for _, book := range books {
		catalogs = catalogs.AppendUnique(book.Catalog)
	}

	catalogs = catalogs.Sort()
	s.cache.Set(KEY, catalogs.Get())
	return catalogs, nil
}
