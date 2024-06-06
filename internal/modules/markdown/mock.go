package markdown

import "github.com/stretchr/testify/mock"

type Mock struct {
	mock.Mock
}

var _ Moduler = (*Mock)(nil)

func (m *Mock) LoadContent(domain string, slug string) (string, error) {
	ret := m.Called(domain, slug)
	return ret.Get(0).(string), ret.Error(1)
}

func (m *Mock) LoadMetadata(domain string, slug string) (map[string]interface{}, error) {
	ret := m.Called(domain, slug)
	return ret.Get(0).(map[string]interface{}), ret.Error(1)
}
