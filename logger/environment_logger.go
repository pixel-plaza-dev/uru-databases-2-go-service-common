package logger

import (
	"log"
)

type EnvironmentLogger struct {
	logger Logger
}

// NewEnvironmentLogger creates a new environment logger
func NewEnvironmentLogger(name string) EnvironmentLogger {
	return EnvironmentLogger{logger: Logger{name}}
}

// ErrorLoadingEnvironmentVariables logs an error message when the environment variables fail to load
func (e *EnvironmentLogger) ErrorLoadingEnvironmentVariables(err error) {
	message := e.logger.buildErrorMessage("Error loading environment variables", err)
	log.Fatalf(message)
}

// VariableNotFound logs an error message when a variable is not found
func (e *EnvironmentLogger) VariableNotFound(fieldName string, variable string) {
	message := fieldName + " not found"
	message = e.logger.buildMessageWithDetails(message, variable)

	log.Fatalf(message)
}
