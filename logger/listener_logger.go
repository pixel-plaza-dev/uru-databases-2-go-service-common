package logger

type ListenerLogger struct {
	logger *Logger
}

// NewListenerLogger creates a new listener logger
func NewListenerLogger(name string) *ListenerLogger {
	return &ListenerLogger{logger: NewLogger(name)}
}

// ServerStarted logs a success message when the server starts
func (l *ListenerLogger) ServerStarted(port string) {
	l.logger.LogMessageWithDetails("Server started", port)
}
