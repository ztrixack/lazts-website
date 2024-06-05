package markdown

import "github.com/stretchr/testify/mock"

type Mock struct {
	mock.Mock
}

var _ Moduler = (*Mock)(nil)

func (m *Mock) ToHTML(filepath string) (string, error) {
	ret := m.Called(filepath)
	return ret.Get(0).(string), ret.Error(1)
}

func (m *Mock) ToMetadata(filepath string) (map[string]interface{}, error) {
	ret := m.Called(filepath)
	return ret.Get(0).(map[string]interface{}), ret.Error(1)
}
