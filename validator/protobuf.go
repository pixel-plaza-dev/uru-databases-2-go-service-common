package validator

import (
	"github.com/pixel-plaza-dev/uru-databases-2-go-service-common/custom_error/protobuf"
	"github.com/pixel-plaza-dev/uru-databases-2-go-service-common/custom_error/validator"
)

// ValidStringFields validates if the string fields are empty
func ValidStringFields(validations *map[string][]error, fields *map[string]string) error {
	failed := false

	// Check if the fields are empty
	for field, value := range *fields {
		if value == "" {
			(*validations)[field] = append((*validations)[field], protobuf.StringIsEmptyError{})
			failed = true
		}
	}

	// If there are no errors, return nil
	if failed {
		return validator.FailedValidationError{}
	}
	return nil
}
