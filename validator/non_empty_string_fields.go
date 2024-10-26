package validator

import (
	"github.com/pixel-plaza-dev/uru-databases-2-go-service-common/custom_error_response/validator"
	"reflect"
)

// ValidNonEmptyStringFields validates if the string fields are empty
func ValidNonEmptyStringFields(validations *map[string][]error, data interface{}, fields *[]string) {
	// Reflection of data
	dataReflection := reflect.ValueOf(data)

	// If data is a pointer, dereference it
	if dataReflection.Kind() == reflect.Ptr {
		dataReflection = dataReflection.Elem()
	}

	// Iterate over the fields
	for _, fieldName := range *fields {
		// Get field
		field := dataReflection.FieldByName(fieldName)

		// Dereference the field if it is a pointer
		if field.Kind() == reflect.Ptr {
			field = field.Elem()
		}

		// Check if the field is a string and is empty
		if field.Kind() == reflect.String && len(field.String()) == 0 {
			(*validations)[fieldName] = append((*validations)[fieldName], validator.StringIsEmptyError{Field: fieldName})
		}
	}
}
