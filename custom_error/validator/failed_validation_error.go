package validator

type FailedValidationError struct{}

// Error returns a formatted error message for FailedValidationError
func (f FailedValidationError) Error() (message string) {
	return "Validation failed"
}
