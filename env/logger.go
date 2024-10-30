package env

import "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/logger"

type Logger struct {
	logger *logger.Logger
}

// NewLogger creates a new environment logger
func NewLogger(name string) *Logger {
	return &Logger{logger: logger.NewLogger(name)}
}

// EnvironmentVariableLoaded logs a message when an environment variable is loaded
func (e *Logger) EnvironmentVariableLoaded(variableName string) {
	e.logger.LogMessageWithDetails("Environment variable loaded", variableName)
}