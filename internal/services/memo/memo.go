package memo

import (
	"lazts/internal/models"
	"lazts/internal/modules/cache"
	"lazts/internal/modules/markdown"
)

type Servicer interface {
	Get(offset uint, limit uint, tag string) ([]models.Memo, error)
	GetTags() ([]models.Tag, error)
	GetSize() int
}

type service struct {
	config     *config
	cache      cache.Moduler
	markdowner markdown.Moduler
}

var _ Servicer = (*service)(nil)

func New(mm markdown.Moduler) *service {
	return &service{
		config:     parseConfig(),
		cache:      cache.New(),
		markdowner: mm,
	}
}
