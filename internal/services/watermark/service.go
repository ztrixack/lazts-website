package watermark

import (
	"image"
	"lazts/internal/modules/imaging"
	"sync"
)

type Servicer interface {
	LoadImage(path string) (image.Image, error)
}

type service struct {
	config *config
	imager imaging.Moduler
	cache  map[string]image.Image
	mutex  *sync.Mutex
}

var _ Servicer = (*service)(nil)

func New(imager imaging.Moduler) *service {
	return &service{
		config: parseConfig(),
		imager: imager,
		cache:  make(map[string]image.Image),
		mutex:  &sync.Mutex{},
	}
}
