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
	if cache, ok := s.cache[KEY].([]models.Tag); ok && cache != nil {
		return cache, nil
	}

	memos, err := s.Get(0, math.MaxInt, "")
	if err != nil {
		return nil, err
	}

	uniqueSet := make(map[string]*models.Tag)
	tags := make([]models.Tag, 0)

	for _, memo := range memos {
		for _, tag := range memo.Tags {
			if existingTag, exists := uniqueSet[tag.Name]; exists {
				existingTag.Count++
			} else {
				tag.Count = 1
				tag.Link = filepath.Join("/memos", strings.ToLower(tag.Name))
				tags = append(tags, tag)
				uniqueSet[tag.Name] = &tags[len(tags)-1]
			}
		}
	}

	sort.Sort(models.TagSort(tags))

	return tags, nil
}
