package issuer

import "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/logger"

type Logger struct {
	logger *logger.Logger
}

// NewLogger creates a new JWT issuer logger
func NewLogger(name string) *Logger {
	return &Logger{logger: logger.NewLogger(name)}
}

// IssuedToken logs a message when the server issues a token
func (e *Logger) IssuedToken() {
	e.logger.LogMessage("Issued token")
}
