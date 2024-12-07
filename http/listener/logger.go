package listener

import (
	commonlogger "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/utils/logger"
)

// Logger is the logger for the listener
type Logger struct {
	logger commonlogger.Logger
}

// NewLogger creates a new listener logger
func NewLogger(logger commonlogger.Logger) (*Logger, error) {
	// Check if the logger is nil
	if logger == nil {
		return nil, commonlogger.NilLoggerError
	}

	return &Logger{logger: logger}, nil
}

// ServerStarted logs a success message when the server starts
func (l Logger) ServerStarted(port string) {
	l.logger.LogMessage(commonlogger.NewLogMessage("Server started", commonlogger.StatusDebug, port))
}
