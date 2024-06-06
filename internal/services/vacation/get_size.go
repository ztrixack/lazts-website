package vacation

func (s *service) GetSize() int {
	if s.size != 0 {
		return s.size
	}

	books, err := s.Get("")
	if err != nil {
		return 0
	}
	s.size = len(books)
	return s.size
}
