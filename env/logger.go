package env

import commonlogger "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/logger"

type Logger struct {
	logger commonlogger.Logger
}

// NewLogger creates a new environment logger
func NewLogger(logger commonlogger.Logger) Logger {
	return Logger{logger: logger}
}

// EnvironmentVariableLoaded logs a message when an environment variable is loaded
func (l Logger) EnvironmentVariableLoaded(variablesName ...string) {
	for _, variableName := range variablesName {
		l.logger.LogMessageWithDetails("Environment variable loaded", variableName)
	}
}
