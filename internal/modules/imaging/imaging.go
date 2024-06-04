package imaging

import (
	"image"
)

type Moduler interface {
	Open(path string) (image.Image, error)
	Resize(img image.Image, width int, height int) *image.NRGBA
	Overlay(background image.Image, img image.Image, pos image.Point, opacity float64) *image.NRGBA
}

type module struct {
}

var _ Moduler = (*module)(nil)

func New() *module {
	return &module{}
}
