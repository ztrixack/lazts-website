package vacation

import (
	"lazts/internal/models"
	"lazts/internal/modules/markdown"
)

type Servicer interface {
	Get(location string) ([]models.Vacation, error)
}

type service struct {
	config     *config
	cache      []models.Vacation
	markdowner markdown.Moduler
}

var _ Servicer = (*service)(nil)

func New(mm markdown.Moduler) *service {
	return &service{
		config:     parseConfig(),
		markdowner: mm,
	}
}
