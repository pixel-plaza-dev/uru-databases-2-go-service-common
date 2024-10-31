package validator

import (
	commonlogger "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/logger"
)

type Logger struct {
	logger commonlogger.Logger
}

// NewLogger creates a new JWT validator logger
func NewLogger(logger commonlogger.Logger) Logger {
	return Logger{logger: logger}
}

// ValidatedToken logs a message when the server validates a token
func (e *Logger) ValidatedToken() {
	e.logger.LogMessage("Validated token")
}
