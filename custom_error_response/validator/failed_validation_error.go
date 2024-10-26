package validator

import "strings"

type FailedValidationError struct {
	FieldsErrors *map[string][]error
}

// Error returns a formatted error message for FailedValidationError
func (f FailedValidationError) Error() string {
	var message strings.Builder
	message.WriteString("Validation failed: { ")

	// Iterate over all fields and their errors
	for field, errors := range *f.FieldsErrors {
		message.WriteString(field)
		message.WriteString(": [")
		for _, err := range errors {
			message.WriteString(err.Error())
			message.WriteString(", ")
		}
		message.WriteString("], ")
	}
	message.WriteString("}")

	return message.String()
}
