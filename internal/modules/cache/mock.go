package cache

import "github.com/stretchr/testify/mock"

type Mock struct {
	mock.Mock
}

func (m *Mock) Get(key string) (interface{}, bool) {
	return nil, false
}

func (m *Mock) LoadAll(keys []string, loader func(string) (interface{}, error)) error {
	return nil
}

func (m *Mock) Set(key string, value interface{}) {
	// noop
}

var _ Moduler = (*Mock)(nil)
