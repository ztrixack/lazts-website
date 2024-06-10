package memo

import (
	"fmt"
	"lazts/internal/models"
	"lazts/internal/utils"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"

	"lazts/internal/modules/log"
)

func (s *service) Get(offset uint, limit uint, tag string) ([]models.Memo, error) {
	KEY := fmt.Sprintf("DATA-%d-%d-%s", offset, limit, tag)
	if value, found := s.cache.Get(KEY); found {
		return value.([]models.Memo), nil
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

	processDir := func(dir os.DirEntry, tag string) {
		defer wg.Done()

		if !dir.IsDir() {
			return
		}

		metadata, err := s.markdowner.LoadMetadata("memos", dir.Name())
		if err != nil {
			errChan <- err
			return
		}

		var memo models.MemoMetadata
		if err := utils.ToStruct(metadata, &memo); err != nil {
			errChan <- err
			return
		}

		mu.Lock()
		defer mu.Unlock()
		if tag != "" && notContains(memo.Tags, tag) {
			log.Fields("dir", dir.Name(), "tag", tag).D("Skipping memo due to tag mismatch")
			return
		}
		if !memo.Published {
			return
		}

		if count < offset {
			count++
			return
		}
		if count >= offset+limit {
			return
		}

		count++
		memos = append(memos, memo.ToMemo())
	}

	for _, dir := range dirs {
		if count >= offset+limit {
			break
		}
		wg.Add(1)
		go processDir(dir, tag)
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	if err := <-errChan; err != nil {
		return nil, err
	}

	sort.Sort(models.MemoSort(memos))
	s.cache.Set(KEY, memos)
	return memos, nil
}

func notContains(tags []string, tag string) bool {
	for _, t := range tags {
		if strings.EqualFold(t, tag) {
			return false
		}
	}
	return true
}
