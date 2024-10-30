package error

import (
	"strings"
)

type FailedValidationError struct {
	FieldsErrors *map[string][]error
}

// Error returns a formatted error message for FailedValidationError
func (f FailedValidationError) Error() string {
	var message strings.Builder
	message.WriteString("Validation failed: { ")

	// Get the number of fields
	fieldsCount := len(*f.FieldsErrors)
	counter := 0

	// Iterate over all fields and their errors
	for field, fieldErrors := range *f.FieldsErrors {
		counter++

		// Add field name
		message.WriteString(field)
		message.WriteString(": [")

		// Iterate over all errors for the field
		for index, err := range fieldErrors {
			message.WriteString(err.Error())

			// Add comma if not the last error
			if index < len(fieldErrors)-1 {
				message.WriteString(", ")
			}
		}

		// Add comma if not the last field
		if counter < fieldsCount {
			message.WriteString("], ")
		} else {
			message.WriteString("]")
		}
	}
	message.WriteString("}")

	return message.String()
}
