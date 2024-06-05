package book

import (
	"lazts/internal/models"
)

func (s *service) GetStats() (*models.BookStats, error) {
	books, err := s.List()
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

	return &models.BookStats{
		Total:     len(books),
		Completed: completed,
		Reading:   reading,
		Unread:    unread,
	}, nil

}
