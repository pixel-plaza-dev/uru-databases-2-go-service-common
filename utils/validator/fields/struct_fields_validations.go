package fields

import (
	"strings"
)

// StructFieldsValidations is a struct that holds the error messages for failed validations of the fields of a struct
type StructFieldsValidations struct {
	FailedFieldsValidations *map[string][]error
	NestedFieldsValidations *map[string]*StructFieldsValidations
}

// NewStructFieldsValidations creates a new StructFieldsValidations struct
func NewStructFieldsValidations() *StructFieldsValidations {
	return &StructFieldsValidations{
		FailedFieldsValidations: &map[string][]error{},
		NestedFieldsValidations: &map[string]*StructFieldsValidations{},
	}
}

// HasFailed returns true if there are failed validations
func (s *StructFieldsValidations) HasFailed() bool {
	// Check if there's a nested failed validation
	if s.NestedFieldsValidations != nil {
		for _, nestedStructFieldsValidations := range *s.NestedFieldsValidations {
			if nestedStructFieldsValidations.HasFailed() {
				return true
			}
		}
	}

	// Check if there are failed fields validations
	if s.FailedFieldsValidations == nil {
		return false
	}
	return len(*s.FailedFieldsValidations) > 0
}

// AddFailedFieldValidationError adds a failed field validation error to the struct
func (s *StructFieldsValidations) AddFailedFieldValidationError(validationName string, validationError error) {
	// Check if the field name is already in the map
	if _, ok := (*s.FailedFieldsValidations)[validationName]; !ok {
		(*s.FailedFieldsValidations)[validationName] = []error{}
	}

	// Append the validation error to the field name
	(*s.FailedFieldsValidations)[validationName] = append((*s.FailedFieldsValidations)[validationName], validationError)
}

// SetNestedFieldsValidations sets the nested struct fields validations to the struct
func (s *StructFieldsValidations) SetNestedFieldsValidations(
	validationName string,
	nestedStructFieldsValidations *StructFieldsValidations,
) {
	(*s.NestedFieldsValidations)[validationName] = nestedStructFieldsValidations
}

// GetLevelPadding returns the padding for the level
func (s *StructFieldsValidations) GetLevelPadding(level int) string {
	var padding strings.Builder
	for i := 0; i < level; i++ {
		padding.WriteString("  ")
	}
	return padding.String()
}

// FailedValidationsMessage returns a formatted error message for StructFieldsValidations
func (s *StructFieldsValidations) FailedValidationsMessage(level int) *string {
	// Check if there are failed validations
	if !s.HasFailed() {
		return nil
	}

	// Get the padding for initial level, the fields, their properties and errors
	basePadding := s.GetLevelPadding(level)
	fieldPadding := s.GetLevelPadding(level + 1)
	fieldPropertiesPadding := s.GetLevelPadding(level + 2)
	fieldErrorsPadding := s.GetLevelPadding(level + 3)

	// Create the message
	var message strings.Builder
	message.WriteString(basePadding)
	message.WriteString("$validations: {\n")

	// Get the number of nested fields validations
	var iteratedFields map[string]bool
	fieldsValidationsNumber := 0
	nestedFieldsValidationsNumber := 0
	iteratedFieldsValidationsNumber := 0
	iteratedNestedFieldsValidationsNumber := 0

	if s.FailedFieldsValidations != nil {
		fieldsValidationsNumber = len(*s.FailedFieldsValidations)
	}
	if s.NestedFieldsValidations != nil {
		nestedFieldsValidationsNumber = len(*s.NestedFieldsValidations)
	}

	// Iterate over all fields and their errors
	for field, fieldErrors := range *s.FailedFieldsValidations {
		iteratedFieldsValidationsNumber++

		// Check if the field has no errors
		if len(fieldErrors) == 0 {
			continue
		}

		// Add field name
		message.WriteString(fieldPadding)
		message.WriteString(field)
		message.WriteString(": {\n")

		// Add field properties flag
		message.WriteString(fieldPropertiesPadding)
		message.WriteString("$errors: [")

		// Iterate over all errors for the field
		iteratedFields[field] = true
		for index, err := range fieldErrors {
			message.WriteString(fieldErrorsPadding)
			message.WriteString(err.Error())

			// Add comma if not the last error
			if index < len(fieldErrors)-1 {
				message.WriteString(",\n")
			} else {
				message.WriteString("\n")
			}
		}

		// Get the nested fields validations for the field if it has any
		var nestedFieldValidations *StructFieldsValidations
		ok := false
		if nestedFieldsValidationsNumber > 0 {
			nestedFieldValidations, ok = (*s.NestedFieldsValidations)[field]
		}

		// Add comma if not it does not have nested fields
		message.WriteString(fieldPropertiesPadding)
		if !ok || !nestedFieldValidations.HasFailed() {
			if ok {
				iteratedNestedFieldsValidationsNumber++
			}

			message.WriteString("]\n")
		} else {
			iteratedNestedFieldsValidationsNumber++
			nestedFieldValidationsMessage := nestedFieldValidations.FailedValidationsMessage(level + 1)

			// Add nested fields errors
			if nestedFieldValidationsMessage != nil {
				message.WriteString("],\n")
				message.WriteString(*nestedFieldValidationsMessage)
			}
		}

		// Add comma if is not the last field
		message.WriteString(fieldPadding)
		if iteratedFieldsValidationsNumber < fieldsValidationsNumber || iteratedNestedFieldsValidationsNumber < nestedFieldsValidationsNumber {
			message.WriteString("},\n")
		} else {
			message.WriteString("}\n")
		}
	}

	// Iterate over all nested fields validations
	if iteratedNestedFieldsValidationsNumber < nestedFieldsValidationsNumber {
		for field, nestedFieldValidations := range *s.NestedFieldsValidations {
			if _, ok := iteratedFields[field]; ok {
				continue
			}

			iteratedNestedFieldsValidationsNumber++
			nestedFieldValidationsMessage := nestedFieldValidations.FailedValidationsMessage(level + 1)

			// Add field name
			message.WriteString(fieldPadding)
			message.WriteString(field)
			message.WriteString(": {\n")

			// Add nested fields errors
			message.WriteString(fieldPropertiesPadding)
			message.WriteString(*nestedFieldValidationsMessage)

			// Add comma if is not the last field
			message.WriteString(fieldPadding)
			if iteratedNestedFieldsValidationsNumber < nestedFieldsValidationsNumber {
				message.WriteString("},\n")
			} else {
				message.WriteString("}\n")
			}
		}
	}

	// Add closing bracket
	message.WriteString(basePadding)
	message.WriteString("}")

	// Get message string
	messageString := message.String()

	return &messageString
}

// String returns a formatted error message f
func (s *StructFieldsValidations) String() string {
	// Return the failed validations message
	message := s.FailedValidationsMessage(0)

	if message != nil {
		return *message
	}
	return ""
}
