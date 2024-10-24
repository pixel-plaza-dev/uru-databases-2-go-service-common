package logger

type MongoDbLogger struct {
	logger *Logger
}

// NewMongoDbLogger is the logger for the MongoDB model
func NewMongoDbLogger(name string) *MongoDbLogger {
	return &MongoDbLogger{logger: NewLogger(name)}
}

// ConnectedToMongoDB logs a success message when the server connects from MongoDB
func (m *MongoDbLogger) ConnectedToMongoDB() {
	m.logger.LogMessage("Connected to MongoDB")
}

// DisconnectedFromMongoDB logs a success message when the server disconnects from MongoDB
func (m *MongoDbLogger) DisconnectedFromMongoDB() {
	m.logger.LogMessage("Disconnected from MongoDB")
}
