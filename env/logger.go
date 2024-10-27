package env

import "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/logger"

type EnvironmentLogger struct {
	logger *logger.Logger
}

// NewEnvironmentLogger creates a new environment logger
func NewEnvironmentLogger(name string) *EnvironmentLogger {

	return &EnvironmentLogger{logger: logger.NewLogger(name)}
}

// EnvironmentVariableLoaded logs a message when an environment variable is loaded
func (e *EnvironmentLogger) EnvironmentVariableLoaded(variableName string) {
	e.logger.LogMessageWithDetails("Environment variable loaded", variableName)
}
