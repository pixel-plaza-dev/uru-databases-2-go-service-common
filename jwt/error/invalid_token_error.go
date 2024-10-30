package error

type InvalidTokenError struct{}

// Error returns the error message
func (e InvalidTokenError) Error() string {
	return "Invalid token"
}
