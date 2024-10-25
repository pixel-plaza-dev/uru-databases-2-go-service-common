package mongodb

type PingToMongoDbFailedError struct {
	Err error
}

// Error returns a formatted error message for PingToMongoDbFailedError
func (l PingToMongoDbFailedError) Error() (message string) {
	return "Failed to ping to MongoDB: " + l.Err.Error()
}
