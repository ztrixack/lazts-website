package vacation

import (
	"lazts/internal/models"
	"lazts/internal/utils"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"
)

func (s *service) Get(location string) ([]models.Vacation, error) {
	if len(s.cache) != 0 {
		return s.cache, nil
	}

	dirs, err := os.ReadDir(filepath.Join(s.config.ContentDir, "vacations"))
	if err != nil {
		return nil, err
	}

	vacations := make([]models.Vacation, 0)
	for _, dir := range dirs {
		if dir.IsDir() {
			metadata, err := s.markdowner.ToMetadata(filepath.Join(s.config.ContentDir, "vacations", dir.Name(), "index.md"))
			if err != nil {
				return nil, err
			}

			var vacation models.VacationMetadata
			if err := utils.ToStruct(metadata, &vacation); err != nil {
				return nil, err
			}

			log.Debug().Interface("vacation", vacation).Str("name", dir.Name()).Msg("vacation on path")

			if location == "" || strings.EqualFold(vacation.Location, location) {
				vacations = append(vacations, vacation.ToHTML())
			}
		}
	}

	s.cache = vacations
	return vacations, nil
}
