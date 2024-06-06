package book

import (
	"lazts/internal/models"
)

func (s *service) GetStats() (*models.BookStats, error) {
	const KEY = "STATS"
	if value, found := s.cache.Get(KEY); found {
		return value.(*models.BookStats), nil
	}

	books, err := s.Get("", "", "")
	if err != nil {
		return nil, err
	}

	completed := 0
	reading := 0
	unread := 0
	for _, book := range books {
		switch book.Status {
		case "completed":
			completed++
		case "reading":
			reading++
		case "unread":
			unread++
		}
	}
	stats := &models.BookStats{
		Total:     len(books),
		Completed: completed,
		Reading:   reading,
		Unread:    unread,
		Pending:   len(books) - completed,
	}

	s.cache.Set(KEY, stats)
	return stats, nil
}
