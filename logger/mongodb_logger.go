package logger

import (
	"log"
)

type MongoDbLogger struct {
	logger Logger
}

// NewMongoDbLogger is the logger for the MongoDB model
func NewMongoDbLogger(name string) MongoDbLogger {
	return MongoDbLogger{logger: Logger{name}}
}

// ErrorConnectingToMongoDB logs an error message when the server fails to connect to MongoDB
func (m *MongoDbLogger) ErrorConnectingToMongoDB(err error) {
	message := m.logger.buildErrorMessage("Error connecting to MongoDB", err)
	log.Fatalf(message)
}
