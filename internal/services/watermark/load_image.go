package watermark

import (
	"image"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

func (s *service) LoadImage(path string) (image.Image, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if img, found := s.cache[path]; found {
		return img, nil
	}

	watermark, err := s.imager.Open(s.config.Path)
	if err != nil {
		log.Error().Err(err).Str("path", s.config.Path).Msg("unable to open watermark")
		return nil, err
	}
	watermark = s.imager.Resize(watermark, s.config.Size, 0)

	imagefile := filepath.Join(s.config.Dir, path)
	original, err := s.imager.Open(imagefile)
	if err != nil {
		log.Error().Err(err).Str("path", path).Msg("unable to open image")
		return nil, err
	}
	offset := image.Pt(original.Bounds().Dx()-watermark.Bounds().Dx()-10, original.Bounds().Dy()-watermark.Bounds().Dy()-10)
	dst := s.imager.Overlay(original, watermark, offset, 1.0)

	s.cache[path] = dst
	return dst, nil
}
