package watermark

import (
	"image"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

func (s *service) LoadImage(imagepath string) (image.Image, error) {
	KEY := "IMAGE-" + imagepath
	if value, found := s.cache.Get(KEY); found {
		return value.(image.Image), nil
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	watermark, err := s.imager.Open(s.config.Path)
	if err != nil {
		log.Error().Err(err).Str("path", s.config.Path).Msg("unable to open watermark")
		return nil, err
	}
	watermark = s.imager.Resize(watermark, s.config.Size, 0)

	imagefile := filepath.Join(s.config.Dir, imagepath)
	original, err := s.imager.Open(imagefile)
	if err != nil {
		log.Error().Err(err).Str("path", imagepath).Msg("unable to open image")
		return nil, err
	}
	offset := image.Pt(original.Bounds().Dx()-watermark.Bounds().Dx()-10, original.Bounds().Dy()-watermark.Bounds().Dy()-10)
	dst := s.imager.Overlay(original, watermark, offset, 1.0)

	s.cache.Set(KEY, dst)
	return dst, nil
}
