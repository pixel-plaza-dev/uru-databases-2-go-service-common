package listener

import "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/logger"

type Logger struct {
	logger *logger.Logger
}

// NewLogger creates a new listener logger
func NewLogger(name string) *Logger {
	return &Logger{logger: logger.NewLogger(name)}
}

// ServerStarted logs a success message when the server starts
func (l *Logger) ServerStarted(port string) {
	l.logger.LogMessageWithDetails("Server started", port)
}
