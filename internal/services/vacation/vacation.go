package vacation

import (
	"lazts/internal/models"
	"lazts/internal/modules/markdown"
)

type Servicer interface {
	Get(location string) ([]models.Vacation, error)
	GetSize() int
}

type service struct {
	config     *config
	cache      map[string][]models.Vacation
	size       int
	markdowner markdown.Moduler
}

var _ Servicer = (*service)(nil)

func New(mm markdown.Moduler) *service {
	return &service{
		config:     parseConfig(),
		cache:      make(map[string][]models.Vacation),
		size:       0,
		markdowner: mm,
	}
}
