package fields

import (
	"reflect"
)

// ValidateNilFields validates if the fields are not nil
func ValidateNilFields(
	data interface{},
	structFieldsToValidate *StructFieldsToValidate,
) (structFieldsValidations *StructFieldsValidations, err error) {
	// Check if either the data or the struct fields to validate are nil
	if data == nil {
		return nil, NilDataError
	}
	if structFieldsToValidate == nil {
		return nil, NilStructFieldsToValidateError
	}

	// Initialize struct fields validations
	structFieldsValidations = NewStructFieldsValidations()

	// Reflection of data
	valueReflection := reflect.ValueOf(data)

	// If data is a pointer, dereference it
	if valueReflection.Kind() == reflect.Ptr {
		valueReflection = valueReflection.Elem()
	}

	// Iterate over the fields
	for fieldName, validationName := range (*structFieldsToValidate).Fields {
		// Get field
		field := valueReflection.FieldByName(fieldName)

		// Check if the field is a pointer
		if field.Kind() == reflect.Ptr {
			// Check if the field exists
			if field.IsNil() {
				structFieldsValidations.AddFailedFieldValidationError(validationName, FieldNotFoundError)
				continue
			}

			// Dereference the field
			field = field.Elem()

			// Check if the field is a nested struct
			if field.Kind() == reflect.Struct {
				// Check if it is a nested struct registered for validation
				nestedStructFieldsToValidate, ok := (*structFieldsToValidate).NestedStructFieldsToValidate[fieldName]
				if !ok {
					continue
				}

				// Validate nested struct
				nestedStructFieldsValidations, err := ValidateNilFields(
					field.Addr().Interface(),
					nestedStructFieldsToValidate,
				)
				if err != nil {
					return nil, err
				}

				// Add nested struct validations to the struct fields validations
				structFieldsValidations.SetNestedFieldsValidations(validationName, nestedStructFieldsValidations)
			}
		}
	}

	return structFieldsValidations, nil
}
