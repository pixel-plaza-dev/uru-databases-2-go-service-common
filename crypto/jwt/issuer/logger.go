package issuer

import (
	commonlogger "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/utils/logger"
)

// Logger is the JWT issuer logger
type Logger struct {
	logger commonlogger.Logger
}

// NewLogger creates a new JWT issuer logger
func NewLogger(logger commonlogger.Logger) Logger {
	return Logger{logger: logger}
}

// IssuedToken logs a message when the server issues a token
func (l *Logger) IssuedToken() {
	l.logger.LogMessage(commonlogger.NewLogMessage("Issued token", commonlogger.StatusInfo))
}
