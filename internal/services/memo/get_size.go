package memo

import "math"

func (s *service) GetSize() int {
	if s.size != 0 {
		return s.size
	}

	books, err := s.Get(0, math.MaxInt)
	if err != nil {
		return 0
	}
	s.size = len(books)
	return s.size
}
