package environment

import (
	"github.com/pixel-plaza-dev/uru-databases-2-go-service-common/utils"
	"strings"
)

type VariableNotFoundError struct {
	VariableName string
}

// Error returns a formatted error message for VariableNotFoundError
func (l VariableNotFoundError) Error() (message string) {
	formattedVariableName := utils.AddBrackets(l.VariableName)
	return strings.Join([]string{"Environment variable not found", formattedVariableName}, " ")
}
