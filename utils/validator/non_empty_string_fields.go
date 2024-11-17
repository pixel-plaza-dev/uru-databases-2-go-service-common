package validator

import (
	"reflect"
)

// ValidNonEmptyStringFields validates if the string fields are empty
func ValidNonEmptyStringFields(validations *map[string][]error, data interface{}, fields *map[string]string) {
	// Reflection of data
	dataReflection := reflect.ValueOf(data)

	// If data is a pointer, dereference it
	if dataReflection.Kind() == reflect.Ptr {
		dataReflection = dataReflection.Elem()
	}

	// Iterate over the fields
	for fieldName, validationName := range *fields {
		// Get field
		field := dataReflection.FieldByName(fieldName)

		// Dereference the field if it is a pointer
		if field.Kind() == reflect.Ptr {
			// Check if the field exists
			if field.IsNil() {
				(*validations)[validationName] = append((*validations)[validationName], FieldNotFoundError)
				continue
			}
			field = field.Elem()
		}

		// Check if the field is a string and is empty
		if field.Kind() == reflect.String && len(field.String()) == 0 {
			(*validations)[validationName] = append((*validations)[validationName], EmptyStringFieldError)
		}
	}
}
