package validator

import (
	"github.com/pixel-plaza-dev/uru-databases-2-go-service-common/logger"
)

type Logger struct {
	logger *logger.Logger
}

// NewLogger creates a new JWT validator logger
func NewLogger(name string) *Logger {
	return &Logger{logger: logger.NewLogger(name)}
}

// ValidatedToken logs a message when the server validates a token
func (e *Logger) ValidatedToken() {
	e.logger.LogMessage("Validated token")
}
