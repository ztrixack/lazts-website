package memo

import (
	"fmt"
	"lazts/internal/models"
	"lazts/internal/utils"
	"os"
	"path/filepath"
	"sort"
	"sync"
)

func (s *service) Get(offset uint, limit uint) ([]models.Memo, error) {
	KEY := fmt.Sprintf("DATA-%d-%d", offset, limit)
	if s.cache[KEY] != nil {
		return s.cache[KEY].([]models.Memo), nil
	}

	dirs, err := os.ReadDir(filepath.Join(s.config.ContentDir, "memos"))
	if err != nil {
		return nil, err
	}

	var (
		memos   []models.Memo
		mu      sync.Mutex
		count   uint
		wg      sync.WaitGroup
		errChan = make(chan error, 1)
	)

	processDir := func(dir os.DirEntry) {
		defer wg.Done()

		if !dir.IsDir() {
			return
		}

		mu.Lock()
		if count < offset {
			count++
			mu.Unlock()
			return
		}
		if count >= offset+limit {
			mu.Unlock()
			return
		}
		count++
		mu.Unlock()

		metadata, err := s.markdowner.ToMetadata(filepath.Join(s.config.ContentDir, "memos", dir.Name(), "index.md"))
		if err != nil {
			errChan <- err
			return
		}

		var metamemo models.MemoMetadata
		if err := utils.ToStruct(metadata, &metamemo); err != nil {
			errChan <- err
			return
		}

		mu.Lock()
		memos = append(memos, metamemo.ToMemo())
		mu.Unlock()
	}

	for _, dir := range dirs {
		if count >= offset+limit {
			break
		}
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

	sort.Sort(models.MemoSort(memos))

	s.cache[KEY] = memos
	return memos, nil
}
