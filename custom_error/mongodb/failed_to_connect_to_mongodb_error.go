package mongodb

type FailedToConnectToMongoDbError struct {
	Err error
}

// Error returns a formatted error message for FailedToConnectToMongoDbError
func (l FailedToConnectToMongoDbError) Error() (message string) {
	return "Failed to connect to MongoDB: " + l.Err.Error()
}
