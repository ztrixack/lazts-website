package imaging

import (
	"image"

	"github.com/stretchr/testify/mock"
)

var _ Moduler = (*Mock)(nil)

type Mock struct {
	mock.Mock
}

func (m *Mock) Open(path string) (image.Image, error) {
	args := m.Called(path)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(image.Image), nil
}

func (m *Mock) Overlay(background image.Image, img image.Image, pos image.Point, opacity float64) *image.NRGBA {
	args := m.Called(background, img, pos, opacity)
	return args.Get(0).(*image.NRGBA)
}

func (m *Mock) Resize(img image.Image, width int, height int) *image.NRGBA {
	args := m.Called(img, width, height)
	return args.Get(0).(*image.NRGBA)
}
