package error

type VariableNotFoundError struct {
	VariableName string
}

// Error returns a formatted error message for VariableNotFoundError
func (l VariableNotFoundError) Error() (message string) {
	return "Environment variable not found: " + l.VariableName
}
