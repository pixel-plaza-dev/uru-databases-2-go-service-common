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
	typeReflection := valueReflection.Type()
	fields := (*structFieldsToValidate).Fields
	nestedStructFieldsToValidate := (*structFieldsToValidate).NestedStructFieldsToValidate
	for i := 0; i < valueReflection.NumField(); i++ {
		fieldValue := valueReflection.Field(i)
		fieldType := typeReflection.Field(i)

		// Check if the field is a pointer
		if fieldValue.Kind() != reflect.Ptr {
			// Check if the field has to be validated
			if fields == nil {
				continue
			}
			validationName, ok := fields[fieldType.Name]
			if !ok {
				continue
			}

			// Check if the field is uninitialized
			if fieldValue.IsZero() {
				structFieldsValidations.AddFailedFieldValidationError(validationName, FieldNotFoundError)
			}
		} else {
			// Check if the field is a nested struct
			if fieldValue.Elem().Kind() != reflect.Struct {
				continue // It's an optional field
			}

			// Check if the nested struct has to be validated
			if fields == nil {
				continue
			}
			validationName, ok := fields[fieldType.Name]
			if !ok {
				continue
			}

			// Check if the field is initialized
			if fieldValue.IsNil() {
				structFieldsValidations.AddFailedFieldValidationError(validationName, FieldNotFoundError)
				continue
			}

			// Get the nested struct fields to validate
			nestedFieldStructFieldsToValidate, ok := nestedStructFieldsToValidate[fieldType.Name]
			if !ok {
				continue
			}

			// Validate nested struct
			nestedStructFieldsValidations, err := ValidateNilFields(
				fieldValue.Addr().Interface(), // TEST IF THIS A POINTER OF THE STRUCT
				nestedFieldStructFieldsToValidate,
			)
			if err != nil {
				return nil, err
			}
			// Add nested struct validations to the struct fields validations
			structFieldsValidations.SetNestedFieldsValidations(validationName, nestedStructFieldsValidations)
		}
	}

	return structFieldsValidations, nil
}
