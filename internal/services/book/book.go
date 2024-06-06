package book

import (
	"lazts/internal/models"
	"lazts/internal/modules/cache"
)

type Servicer interface {
	Get(search, catalog, status string) ([]models.Book, error)
	GetStats() (*models.BookStats, error)
	GetShelf(...int) ([][]models.Book, error)
	GetCatalogs() ([]models.Option, error)
	GetSize() int
}

type service struct {
	config *config
	cache  cache.Moduler
}

var _ Servicer = (*service)(nil)

func New() *service {
	return &service{
		config: parseConfig(),
		cache:  cache.New(),
	}
}
