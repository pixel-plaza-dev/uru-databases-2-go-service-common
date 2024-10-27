package error

type PasswordNotHashedError struct{}

// Error returns a formatted error message for PasswordNotHashedError
func (p PasswordNotHashedError) Error() (message string) {
	return "Password is not hashed"
}
