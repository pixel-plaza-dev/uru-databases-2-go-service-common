package mongodb

import (
	"github.com/pixel-plaza-dev/uru-databases-2-go-service-common/utils"
	"strings"
)

type FailedToStartSessionError struct {
	Err error
}

// Error returns a formatted error message for FailedToStartSessionError
func (f FailedToStartSessionError) Error() (message string) {
	formattedError := utils.AddBrackets(f.Err.Error())
	return strings.Join([]string{"Failed to start session", formattedError}, " ")
}
