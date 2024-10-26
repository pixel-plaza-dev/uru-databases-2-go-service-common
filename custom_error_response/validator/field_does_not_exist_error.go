package validator

type FieldDoesNotExistError struct{}

// Error returns the error message
func (e FieldDoesNotExistError) Error() string {
	return "Field does not exist"
}
