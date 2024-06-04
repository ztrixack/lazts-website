package imaging

import (
	"image"

	"github.com/disintegration/imaging"
)

func (m *module) Resize(img image.Image, width int, height int) *image.NRGBA {
	return imaging.Resize(img, width, height, imaging.Lanczos)
}
