package error

// VariableNotFoundError is the error type for when a variable is not found in the environment
type VariableNotFoundError struct {
	Key string
}

// Error returns the error message
func (e VariableNotFoundError) Error() string {
	return "Variable not found in environment: " + e.Key
}
