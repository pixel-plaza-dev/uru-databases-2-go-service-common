package listener

type FailedToCloseError struct {
	Err error
}

// Error returns a formatted error message for FailedToCloseError
func (l FailedToCloseError) Error() (message string) {
	return "Failed to close listener: " + l.Err.Error()
}
