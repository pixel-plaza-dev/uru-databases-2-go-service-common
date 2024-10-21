package logger

import (
	"github.com/pixel-plaza-dev/uru-databases-2-go-service-common/custom_error/environment"
)

type EnvironmentLogger struct {
	logger Logger
}

// NewEnvironmentLogger creates a new environment logger
func NewEnvironmentLogger(name string) *EnvironmentLogger {
	return &EnvironmentLogger{logger: Logger{name}}
}

// ErrorLoadingEnvironmentVariables logs an custom_error message when the environment variables fail to load
func (e *EnvironmentLogger) ErrorLoadingEnvironmentVariables(err error) {
	environmentError := environment.LoadingEnvironmentVariablesError{Err: err}
	e.logger.logError(environmentError)
}

// VariableNotFound logs an error message when an environment variable is not found
func (e *EnvironmentLogger) VariableNotFound(variable string) {
	variableNotFound := environment.VariableNotFoundError{Variable: variable}
	e.logger.logError(variableNotFound)
}
