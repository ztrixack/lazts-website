package book

func (s *service) GetSize() int {
	const KEY = "SIZE"
	if value, found := s.cache.Get(KEY); found {
		return value.(int)
	}

	books, err := s.Get("", "", "")
	if err != nil {
		return 0
	}

	s.cache.Set(KEY, len(books))
	return len(books)
}
