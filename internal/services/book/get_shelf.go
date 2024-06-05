package book

import (
	"lazts/internal/models"
	"math/rand/v2"
)

func (s *service) GetShelf(rows ...int) ([][]models.Book, error) {
	books, err := s.List()
	if err != nil {
		return nil, err
	}

	result := make([][]models.Book, len(rows))
	for i := range result {
		result[i] = make([]models.Book, rows[i])
		for j := range result[i] {
			r := rand.IntN(len(books))
			result[i][j] = books[r]
		}
	}

	return result, nil
}
