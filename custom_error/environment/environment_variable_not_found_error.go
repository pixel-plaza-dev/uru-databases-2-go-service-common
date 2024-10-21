package environment

type VariableNotFoundError struct {
	Variable string
}

// Error returns a formatted error message for VariableNotFoundError
func (l VariableNotFoundError) Error() (message string) {
	return "Environment variable not found: " + l.Variable
}
