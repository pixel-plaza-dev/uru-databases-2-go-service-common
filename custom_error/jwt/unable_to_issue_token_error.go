package jwt

import (
	"github.com/pixel-plaza-dev/uru-databases-2-go-service-common/utils"
	"strings"
)

type UnableToIssueTokenError struct {
	Err error
}

// Error returns a formatted error message for UnableToIssueTokenError
func (u UnableToIssueTokenError) Error() (message string) {
	formattedError := utils.AddBrackets(u.Err.Error())
	return strings.Join([]string{"Unable to issue token", formattedError}, " ")
}
