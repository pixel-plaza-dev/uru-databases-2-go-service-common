package validator

type StringIsEmptyError struct {
	Field string
}

// Error returns a formatted error message for StringIsEmptyError
func (f StringIsEmptyError) Error() (message string) {
	return "String field is empty: " + f.Field
}
