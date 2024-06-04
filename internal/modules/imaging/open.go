package imaging

import (
	"image"

	"github.com/disintegration/imaging"
)

func (m *module) Open(path string) (image.Image, error) {
	return imaging.Open(path)
}
