package jwt

type UnableToParseTokenError struct {
	Err error
}

// Error returns a formatted error message for UnableToParseTokenError
func (u UnableToParseTokenError) Error() (message string) {
	return "Unable to parse token: " + u.Err.Error()
}
