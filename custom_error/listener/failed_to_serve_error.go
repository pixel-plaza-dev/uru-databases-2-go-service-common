package listener

type FailedToServeError struct {
	Err error
}

// Error returns a formatted error message for FailedToServeError
func (l FailedToServeError) Error() (message string) {
	return "Failed to serve: [" + l.Err.Error() + "]"
}
