package listener

type FailedToListenError struct {
	Err error
}

// Error returns a formatted error message for FailedToListenError
func (l FailedToListenError) Error() (message string) {
	return "Failed to listen: " + l.Err.Error()
}
