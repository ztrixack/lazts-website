package log

var logger = New(Config()).logger

func D(format string, args ...interface{}) {
	logger.Debug().CallerSkipFrame(1).Msgf(format, args...)
}

func I(format string, args ...interface{}) {
	logger.Info().CallerSkipFrame(1).Msgf(format, args...)
}

func W(format string, args ...interface{}) {
	logger.Warn().CallerSkipFrame(1).Msgf(format, args...)
}

func E(format string, args ...interface{}) {
	logger.Error().CallerSkipFrame(1).Msgf(format, args...)
}

func C(format string, args ...interface{}) {
	logger.Fatal().CallerSkipFrame(1).Msgf(format, args...)
}

func Fields(fields ...interface{}) Moduler {
	newLogger := logger.With()
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
		logger: &logger,
	}
}

func Err(err error) Moduler {
	logger := logger.With().Err(err).Logger()

	return &module{
		logger: &logger,
	}
}
