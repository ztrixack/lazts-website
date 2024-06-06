package markdown

import "github.com/stretchr/testify/mock"

type Mock struct {
	mock.Mock
}

var _ Moduler = (*Mock)(nil)

func (m *Mock) ReadFile(name string, filepath string) ([]byte, error) {
	ret := m.Called(name, filepath)
	return ret.Get(0).([]byte), ret.Error(1)
}

func (m *Mock) ToHTML(name string, data []byte) (string, error) {
	ret := m.Called(name, data)
	return ret.Get(0).(string), ret.Error(1)
}

func (m *Mock) ToMetadata(name string, data []byte) (map[string]interface{}, error) {
	ret := m.Called(name, data)
	return ret.Get(0).(map[string]interface{}), ret.Error(1)
}
