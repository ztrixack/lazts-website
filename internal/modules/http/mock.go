package http

import (
	"context"
	"net/http"

	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
}

var _ Moduler = (*Mock)(nil)

func (m *Mock) Serve() error {
	ret := m.Called()
	return ret.Error(0)
}

func (m *Mock) Shutdown(c context.Context) error {
	ret := m.Called(c)
	return ret.Error(0)
}

func (m *Mock) Get(path string, handler http.HandlerFunc) {
	m.Called(path, handler)
}

func (m *Mock) Use(middlewares ...func(http.Handler) http.Handler) {
	m.Called(middlewares)
}
