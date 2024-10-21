package environment

type LoadingEnvironmentVariablesError struct {
	Err error
}

// Error returns a formatted error message for LoadingEnvironmentVariablesError
func (l LoadingEnvironmentVariablesError) Error() (message string) {
	return "Error loading environment variables: [" + l.Err.Error() + "]"
}
