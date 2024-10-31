package listener

import commonlogger "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/logger"

// Logger is the logger for the listener
type Logger struct {
	logger commonlogger.Logger
}

// NewLogger creates a new listener logger
func NewLogger(logger commonlogger.Logger) Logger {
	return Logger{logger: logger}
}

// ServerStarted logs a success message when the server starts
func (l *Logger) ServerStarted(port string) {
	l.logger.LogMessageWithDetails("Server started", port)
}
