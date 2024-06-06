package log

import (
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/mock"
)

var _ Moduler = (*Mock)(nil)

type Mock struct {
	mock.Mock
}

func (l *Mock) Config() config {
	return config{}
}

func (l *Mock) D(format string, args ...interface{}) {
	// l.Called(append([]interface{}{format}, args...)...)
}

func (l *Mock) I(format string, args ...interface{}) {
	// l.Called(append([]interface{}{format}, args...)...)
}

func (l *Mock) W(format string, args ...interface{}) {
	// l.Called(append([]interface{}{format}, args...)...)
}

func (l *Mock) E(format string, args ...interface{}) {
	// l.Called(append([]interface{}{format}, args...)...)
}

func (l *Mock) C(format string, args ...interface{}) {
	// l.Called(append([]interface{}{format}, args...)...)
}

func (l *Mock) Fields(fields ...interface{}) Moduler {
	// l.Called(fields...)
	return l
}

func (l *Mock) Err(err error) Moduler {
	// l.Called(err)
	return l
}

func (l *Mock) GetLogger() zerolog.Logger {
	// l.Called(key, value)
	return zerolog.Logger{}
}
