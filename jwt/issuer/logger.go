package issuer

import commonlogger "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/logger"

type Logger struct {
	logger commonlogger.Logger
}

// NewLogger creates a new JWT issuer logger
func NewLogger(logger commonlogger.Logger) Logger {
	return Logger{logger: logger}
}

// IssuedToken logs a message when the server issues a token
func (e *Logger) IssuedToken() {
	e.logger.LogMessage("Issued token")
}
