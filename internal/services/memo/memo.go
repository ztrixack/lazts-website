package memo

import (
	"lazts/internal/models"
	"lazts/internal/modules/markdown"
)

type Servicer interface {
	Get(limit uint, offset uint) ([]models.Memo, error)
	GetSize() int
}

type service struct {
	config     *config
	cache      map[string][]models.Memo
	size       int
	markdowner markdown.Moduler
}

var _ Servicer = (*service)(nil)

func New(mm markdown.Moduler) *service {
	return &service{
		config:     parseConfig(),
		cache:      make(map[string][]models.Memo),
		size:       0,
		markdowner: mm,
	}
}
