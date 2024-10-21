package logger

import "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/custom_error/mongodb"

type MongoDbLogger struct {
	logger Logger
}

// NewMongoDbLogger is the logger for the MongoDB model
func NewMongoDbLogger(name string) *MongoDbLogger {
	return &MongoDbLogger{logger: Logger{name}}
}

// ConnectedToMongoDB logs a success message when the server connects from MongoDB
func (m *MongoDbLogger) ConnectedToMongoDB() {
	message := m.logger.buildMessage("Connected to MongoDB")
	m.logger.logMessage(message)
}

// FailedToConnectToMongoDb logs an error message when the server fails to connect to MongoDB
func (m *MongoDbLogger) FailedToConnectToMongoDb(err error) {
	failedToConnectToMongoDb := mongodb.FailedToConnectToMongoDbError{Err: err}
	m.logger.logError(failedToConnectToMongoDb)
}

// DisconnectedFromMongoDB logs a success message when the server disconnects from MongoDB
func (m *MongoDbLogger) DisconnectedFromMongoDB() {
	message := m.logger.buildMessage("Disconnected from MongoDB")
	m.logger.logMessage(message)
}

// FailedToDisconnectFromMongoDb logs an error message when the server fails to disconnect from MongoDB
func (m *MongoDbLogger) FailedToDisconnectFromMongoDb(err error) {
	failedToDisconnectFromMongoDb := mongodb.FailedToDisconnectFromMongoDbError{Err: err}
	m.logger.logError(failedToDisconnectFromMongoDb)
}
