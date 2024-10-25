package mongodb

type FailedToStartSessionError struct {
	Err error
}

// Error returns a formatted error message for FailedToStartSessionError
func (f FailedToStartSessionError) Error() (message string) {
	return "Failed to start session: " + f.Err.Error()
}
