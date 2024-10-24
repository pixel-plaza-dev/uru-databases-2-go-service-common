package validator

import (
	"github.com/pixel-plaza-dev/uru-databases-2-go-service-common/custom_error/protobuf"
)

// ValidStringFields validates if the string fields are empty
func ValidStringFields(validations *map[string][]error, fields *map[string]string) {
	for field, value := range *fields {
		if value == "" {
			(*validations)[field] = append((*validations)[field], protobuf.StringIsEmptyError{})
		}
	}
}
