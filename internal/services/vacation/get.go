package vacation

import (
	"lazts/internal/models"
	"lazts/internal/utils"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
)

func (s *service) Get(location string) ([]models.Vacation, error) {
	cpath := location
	if s.cache[cpath] != nil {
		return s.cache[cpath], nil
	}

	dirs, err := os.ReadDir(filepath.Join(s.config.ContentDir, "vacations"))
	if err != nil {
		return nil, err
	}

	var (
		vacations []models.Vacation
		wg        sync.WaitGroup
		errChan   = make(chan error, 1)
	)

	processDir := func(dir os.DirEntry) {
		defer wg.Done()

		if !dir.IsDir() {
			return
		}

		metadata, err := s.markdowner.ToMetadata(filepath.Join(s.config.ContentDir, "vacations", dir.Name(), "index.md"))
		if err != nil {
			errChan <- err
			return
		}

		var vacation models.VacationMetadata
		if err := utils.ToStruct(metadata, &vacation); err != nil {
			errChan <- err
			return
		}

		if location == "" || strings.EqualFold(vacation.Location, location) {
			vacations = append(vacations, vacation.ToHTML())
		}
	}

	for _, dir := range dirs {
		wg.Add(1)
		go processDir(dir)
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	if err := <-errChan; err != nil {
		return nil, err
	}

	sort.Sort(models.VacationSort(vacations))

	s.cache[cpath] = vacations
	return vacations, nil
}
