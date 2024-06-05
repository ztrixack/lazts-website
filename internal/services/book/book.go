package book

import "lazts/internal/models"

type Servicer interface {
	Get(search, catalog, status string) ([]models.Book, error)
	GetStats() (*models.BookStats, error)
	GetShelf(...int) ([][]models.Book, error)
	GetCatalogs() ([]models.Option, error)
	GetSize() int
}

type service struct {
	config   *config
	caches   map[string][]models.Book
	catalogs []models.Option
	size     int
}

var _ Servicer = (*service)(nil)

func New() *service {
	return &service{
		config:   parseConfig(),
		caches:   make(map[string][]models.Book),
		catalogs: nil,
		size:     0,
	}
}
