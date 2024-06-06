package memo

import "math"

func (s *service) GetSize() int {
	const KEY = "SIZE"
	if cache, ok := s.cache[KEY].(int); ok && cache != 0 {
		return cache
	}

	books, err := s.Get(0, math.MaxInt, "")
	if err != nil {
		return 0
	}
	s.cache[KEY] = len(books)
	return len(books)
}
