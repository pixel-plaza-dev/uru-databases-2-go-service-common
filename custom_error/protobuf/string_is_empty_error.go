package protobuf

import (
	"github.com/pixel-plaza-dev/uru-databases-2-go-service-common/utils"
	"strings"
)

type StringIsEmptyError struct {
	Field string
}

// Error returns a formatted error message for StringIsEmptyError
func (f StringIsEmptyError) Error() (message string) {
	formattedField := utils.AddBrackets(f.Field)
	return strings.Join([]string{"String field is empty", formattedField}, " ")
}
