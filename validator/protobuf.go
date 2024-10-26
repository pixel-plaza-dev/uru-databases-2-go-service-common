package validator

import (
	"github.com/pixel-plaza-dev/uru-databases-2-go-service-common/custom_error_response/validator"
	"reflect"
)

// ValidNonEmptyStringFields validates if the string fields are empty
func ValidNonEmptyStringFields(validations *map[string][]error, data interface{}, fields *[]string) {
	// Reflection of data
	dataReflection := reflect.ValueOf(data)

	// Iterate over the fields
	for _, fieldName := range *fields {
		// Get field
		field := dataReflection.FieldByName(fieldName).String()

		if len(field) == 0 {
			(*validations)[fieldName] = append((*validations)[fieldName], validator.StringIsEmptyError{Field: fieldName})
		}
	}
}
