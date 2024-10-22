package jwt

import (
	"github.com/pixel-plaza-dev/uru-databases-2-go-service-common/utils"
	"strings"
)

type UnableToParseTokenError struct {
	Err error
}

// Error returns a formatted error message for UnableToParseTokenError
func (u UnableToParseTokenError) Error() (message string) {
	formattedError := utils.AddBrackets(u.Err.Error())
	return strings.Join([]string{"Unable to parse token string", formattedError}, " ")
}
