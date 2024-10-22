package logger

type EnvironmentLogger struct {
	logger *Logger
}

// NewEnvironmentLogger creates a new environment logger
func NewEnvironmentLogger(name string) *EnvironmentLogger {

	return &EnvironmentLogger{logger: NewLogger(name)}
}

// EnvironmentVariableLoaded logs a message when an environment variable is loaded
func (e *EnvironmentLogger) EnvironmentVariableLoaded(variableName string) {
	e.logger.logMessageWithDetails("Environment variable loaded", variableName)
}
