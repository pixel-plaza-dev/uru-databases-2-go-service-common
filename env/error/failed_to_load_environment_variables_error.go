package error

type FailedToLoadEnvironmentVariablesError struct {
	Err error
}

// Error returns a formatted error message for FailedToLoadEnvironmentVariablesError
func (l FailedToLoadEnvironmentVariablesError) Error() (message string) {
	return "Failed to load environment variables: " + l.Err.Error()
}
