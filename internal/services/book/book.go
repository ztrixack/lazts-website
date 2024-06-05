package book

import "lazts/internal/models"

type Servicer interface {
	List() ([]models.Book, error)
	GetStats() (*models.BookStats, error)
	GetShelf(...int) ([][]models.Book, error)
}

type service struct {
	config *config
	cache  []models.Book
}

var _ Servicer = (*service)(nil)

func New() *service {
	return &service{
		config: parseConfig(),
	}
}
