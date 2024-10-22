package listener

import (
	"github.com/pixel-plaza-dev/uru-databases-2-go-service-common/utils"
	"strings"
)

type FailedToListenError struct {
	Err error
}

// Error returns a formatted error message for FailedToListenError
func (l FailedToListenError) Error() (message string) {
	formattedError := utils.AddBrackets(l.Err.Error())
	return strings.Join([]string{"Failed to listen", formattedError}, " ")
}
