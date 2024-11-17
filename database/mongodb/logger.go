package mongodb

import (
	commonlogger "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/utils/logger"
)

// Logger is the logger for the MongoDB model
type Logger struct {
	logger commonlogger.Logger
}

// NewLogger is the logger for the MongoDB model
func NewLogger(logger commonlogger.Logger) Logger {
	return Logger{logger: logger}
}

// ConnectedToMongoDB logs a success message when the server connects from MongoDB
func (l Logger) ConnectedToMongoDB() {
	l.logger.LogMessage("Connected to MongoDB")
}

// DisconnectedFromMongoDB logs a success message when the server disconnects from MongoDB
func (l Logger) DisconnectedFromMongoDB() {
	l.logger.LogMessage("Disconnected from MongoDB")
}
