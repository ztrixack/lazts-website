package web

import (
	"io"

	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
}

var _ Servicer = (*Mock)(nil)

func (m *Mock) RenderMarkdown(w io.Writer, path string, data map[string]interface{}) error {
	ret := m.Called(w, path, data)
	return ret.Error(0)
}

func (m *Mock) RenderPartial(w io.Writer, path string, data map[string]interface{}) error {
	ret := m.Called(w, path, data)
	return ret.Error(0)
}

func (m *Mock) RenderPage(w io.Writer, path string, data map[string]interface{}) error {
	ret := m.Called(w, path, data)
	return ret.Error(0)
}
