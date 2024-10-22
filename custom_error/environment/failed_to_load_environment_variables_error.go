package environment

import (
	"github.com/pixel-plaza-dev/uru-databases-2-go-service-common/utils"
	"strings"
)

type FailedToLoadEnvironmentVariablesError struct {
	Err error
}

// Error returns a formatted error message for FailedToLoadEnvironmentVariablesError
func (l FailedToLoadEnvironmentVariablesError) Error() (message string) {
	formattedError := utils.AddBrackets(l.Err.Error())
	return strings.Join([]string{"Failed to load environment variables", formattedError}, " ")
}
