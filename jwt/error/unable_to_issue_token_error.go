package error

type UnableToIssueTokenError struct {
	Err error
}

// Error returns a formatted error message for UnableToIssueTokenError
func (u UnableToIssueTokenError) Error() (message string) {
	return "Unable to issue token: " + u.Err.Error()
}
