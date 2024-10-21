package logger

import "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/custom_error/listener"

type ListenerLogger struct {
	logger Logger
}

// NewListenerLogger creates a new listener logger
func NewListenerLogger(name string) *ListenerLogger {
	return &ListenerLogger{logger: Logger{name}}
}

// FailedToListen logs an error message when the server fails to listen
func (l *ListenerLogger) FailedToListen(err error) {
	failedToListen := listener.FailedToListenError{Err: err}
	l.logger.logError(failedToListen)
}

// ServerStarted logs a success message when the server starts
func (l *ListenerLogger) ServerStarted(port string) {
	message := l.logger.buildMessage("Server started on: " + port)
	l.logger.logMessage(message)
}

// FailedToServe logs an error message when the server fails to serve
func (l *ListenerLogger) FailedToServe(err error) {
	failedToServe := listener.FailedToServeError{Err: err}
	l.logger.logError(failedToServe)
}

// FailedToClose logs an error message when the server fails to close
func (l *ListenerLogger) FailedToClose(err error) {
	failedToClose := listener.FailedToCloseError{Err: err}
	l.logger.logError(failedToClose)
}
