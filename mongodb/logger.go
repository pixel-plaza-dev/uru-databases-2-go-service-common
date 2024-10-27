package mongodb

import "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/logger"

type Logger struct {
	logger *logger.Logger
}

// NewLogger is the logger for the MongoDB model
func NewLogger(name string) *Logger {
	return &Logger{logger: logger.NewLogger(name)}
}

// ConnectedToMongoDB logs a success message when the server connects from MongoDB
func (m *Logger) ConnectedToMongoDB() {
	m.logger.LogMessage("Connected to MongoDB")
}

// DisconnectedFromMongoDB logs a success message when the server disconnects from MongoDB
func (m *Logger) DisconnectedFromMongoDB() {
	m.logger.LogMessage("Disconnected from MongoDB")
}
