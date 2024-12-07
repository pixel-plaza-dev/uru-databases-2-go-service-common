package env

import (
	commonlogger "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/utils/logger"
)

type Logger struct {
	logger commonlogger.Logger
}

// NewLogger creates a new environment logger
func NewLogger(logger commonlogger.Logger) (*Logger, error) {
	// Check if the logger is nil
	if logger == nil {
		return nil, commonlogger.NilLoggerError
	}

	return &Logger{logger: logger}, nil
}

// EnvironmentVariableLoaded logs a message when an environment variable is loaded
func (l Logger) EnvironmentVariableLoaded(variablesName ...string) {
	l.logger.LogMessage(
		commonlogger.NewLogMessage(
			"Environment variable loaded",
			commonlogger.StatusDebug,
			variablesName...,
		),
	)
}
