package jwt

type UnexpectedSigningMethodError struct {
	Algorithm interface{}
}

// Error returns a formatted error message for UnexpectedSigningMethodError
func (u UnexpectedSigningMethodError) Error() (message string) {
	return "Unexpected signing method: " + u.Algorithm.(string)
}
