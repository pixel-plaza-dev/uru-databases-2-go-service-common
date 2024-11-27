package env

import (
	commonlogger "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/utils/logger"
)

type Logger struct {
	logger commonlogger.Logger
}

// NewLogger creates a new environment logger
func NewLogger(logger commonlogger.Logger) Logger {
	return Logger{logger: logger}
}

// EnvironmentVariableLoaded logs a message when an environment variable is loaded
func (l Logger) EnvironmentVariableLoaded(variablesName ...string) {
	l.logger.LogMessage(commonlogger.NewLogMessage("Environment variable loaded", commonlogger.StatusDebug, variablesName...))
}
