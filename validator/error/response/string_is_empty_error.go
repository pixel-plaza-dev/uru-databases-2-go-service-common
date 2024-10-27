package response

type StringIsEmptyError struct{}

// Error returns a formatted error message for StringIsEmptyError
func (f StringIsEmptyError) Error() (message string) {
	return "String field is empty"
}
