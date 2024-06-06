package memo

import (
	"lazts/internal/models"
	"lazts/internal/modules/markdown"
)

type Servicer interface {
	Get(offset uint, limit uint, tag string) ([]models.Memo, error)
	GetTags() ([]models.Tag, error)
	GetSize() int
}

type service struct {
	config     *config
	cache      map[string]interface{}
	markdowner markdown.Moduler
}

var _ Servicer = (*service)(nil)

func New(mm markdown.Moduler) *service {
	return &service{
		config:     parseConfig(),
		cache:      make(map[string]interface{}),
		markdowner: mm,
	}
}
