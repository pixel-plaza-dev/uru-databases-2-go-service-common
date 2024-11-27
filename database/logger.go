package database

import (
	commonlogger "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/utils/logger"
)

// Logger is the logger for the database connection
type Logger struct {
	logger commonlogger.Logger
}

// NewLogger is the logger for the database connection
func NewLogger(logger commonlogger.Logger) Logger {
	return Logger{logger: logger}
}

// ConnectedToDatabase logs a success message when the server connects to the database
func (l Logger) ConnectedToDatabase() {
	l.logger.LogMessage(commonlogger.NewLogMessage("Connected to database", commonlogger.StatusDebug))
}

// DisconnectedFromDatabase logs a success message when the server disconnects from the database
func (l Logger) DisconnectedFromDatabase() {
	l.logger.LogMessage(commonlogger.NewLogMessage("Disconnected from database", commonlogger.StatusDebug))
}
