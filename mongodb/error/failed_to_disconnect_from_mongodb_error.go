package error

type FailedToDisconnectFromMongoDbError struct {
	Err error
}

// Error returns a formatted error message for FailedToDisconnectFromMongoDbError
func (l FailedToDisconnectFromMongoDbError) Error() (message string) {
	return "Failed to disconnect from MongoDB: " + l.Err.Error()
}
