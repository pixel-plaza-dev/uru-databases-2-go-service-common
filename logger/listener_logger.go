package logger

import (
	"log"
)

type ListenerLogger struct {
	logger Logger
}

// NewListenerLogger creates a new listener logger
func NewListenerLogger(name string) ListenerLogger {
	return ListenerLogger{logger: Logger{name}}
}

// FailedToListen logs an error message when the grpc_server fails to listen
func (l *ListenerLogger) FailedToListen(err error) {
	message := l.logger.buildErrorMessage("Failed to listen", err)
	log.Fatalf(message)
}

// ServerStarted logs a success message when the grpc_server starts
func (l *ListenerLogger) ServerStarted(port string) {
	message := l.logger.buildMessage("Server started on: " + port)
	log.Println(message)
}

// FailedToServe logs an error message when the grpc_server fails to serve
func (l *ListenerLogger) FailedToServe(err error) {
	message := l.logger.buildErrorMessage("Failed to serve", err)
	log.Fatalf(message)
}
