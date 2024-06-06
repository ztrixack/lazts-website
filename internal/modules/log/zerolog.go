package log

import (
	"time"

	"github.com/rs/zerolog"
)

type Moduler interface {
	D(format string, args ...interface{})
	I(format string, args ...interface{})
	W(format string, args ...interface{})
	E(format string, args ...interface{})
	C(format string, args ...interface{})
	Fields(kv ...interface{}) Moduler
	Err(err error) Moduler
}

type module struct {
	config *config
	logger *zerolog.Logger
}

var _ Moduler = (*module)(nil)

func New(c *config) *module {
	writer := zerolog.ConsoleWriter{
		Out:        c.Writer,
		NoColor:    true,
		TimeFormat: time.RFC3339,
	}

	logger := zerolog.New(writer).
		Level(zerolog.Level(c.Level)).
		With().
		Timestamp().
		Caller().
		Logger()

	return &module{
		config: c,
		logger: &logger,
	}
}

func (l *module) D(format string, args ...interface{}) {
	l.logger.Debug().CallerSkipFrame(1).Msgf(format, args...)
}

func (l *module) I(format string, args ...interface{}) {
	l.logger.Info().CallerSkipFrame(1).Msgf(format, args...)
}

func (l *module) W(format string, args ...interface{}) {
	l.logger.Warn().CallerSkipFrame(1).Msgf(format, args...)
}

func (l *module) E(format string, args ...interface{}) {
	l.logger.Error().CallerSkipFrame(1).Msgf(format, args...)
}

func (l *module) C(format string, args ...interface{}) {
	l.logger.Fatal().CallerSkipFrame(1).Msgf(format, args...)
}

func (l *module) Fields(fields ...interface{}) Moduler {
	newLogger := l.logger.With()
	k := ""
	for i, v := range fields {
		if i%2 == 0 {
			k = v.(string)
			continue
		}
		newLogger = newLogger.Interface(k, v)
	}
	logger := newLogger.Logger()

	return &module{
		config: l.config,
		logger: &logger,
	}
}

func (l *module) Err(err error) Moduler {
	logger := l.logger.With().Err(err).Logger()

	return &module{
		config: l.config,
		logger: &logger,
	}
}
