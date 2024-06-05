package memo

import (
	"lazts/internal/models"
	"lazts/internal/modules/markdown"
)

type Servicer interface {
	Get(limit uint, offset uint) ([]models.Memo, error)
}

type service struct {
	config     *config
	cache      []models.Memo
	markdowner markdown.Moduler
}

var _ Servicer = (*service)(nil)

func New(mm markdown.Moduler) *service {
	return &service{
		config:     parseConfig(),
		markdowner: mm,
	}
}
