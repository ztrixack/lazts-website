package memo

import (
	"lazts/internal/models"
	"math"
	"path/filepath"
	"sort"
	"strings"
)

func (s *service) GetTags() ([]models.Tag, error) {
	const KEY = "TAGS"
	if value, found := s.cache.Get(KEY); found {
		return value.([]models.Tag), nil
	}

	memos, err := s.Get(0, math.MaxInt, "")
	if err != nil {
		return nil, err
	}

	uniqueSet := make(map[string]struct{})
	tags := make([]models.Tag, 0)

	for _, memo := range memos {
		for _, tag := range memo.Tags {
			if _, exists := uniqueSet[tag.Name]; exists {
				for i, t := range tags {
					if t.Name == tag.Name {
						tags[i].Count++
					}
				}
			} else {
				tag.Count = 1
				tag.Link = filepath.Join("/memos", strings.ToLower(tag.Name))
				tags = append(tags, tag)
				uniqueSet[tag.Name] = struct{}{}
			}
		}
	}

	sort.Sort(models.TagSort(tags))
	s.cache.Set(KEY, tags)
	return tags, nil
}
