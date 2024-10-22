package listener

import (
	"github.com/pixel-plaza-dev/uru-databases-2-go-service-common/utils"
	"strings"
)

type FailedToCloseError struct {
	Err error
}

// Error returns a formatted error message for FailedToCloseError
func (l FailedToCloseError) Error() (message string) {
	formattedError := utils.AddBrackets(l.Err.Error())
	return strings.Join([]string{"Failed to close", formattedError}, " ")
}
