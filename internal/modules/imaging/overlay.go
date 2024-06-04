package imaging

import (
	"image"

	"github.com/disintegration/imaging"
)

func (m *module) Overlay(background image.Image, img image.Image, pos image.Point, opacity float64) *image.NRGBA {
	return imaging.Overlay(background, img, pos, opacity)
}
