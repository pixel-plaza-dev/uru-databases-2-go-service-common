package jwt

import (
	"github.com/pixel-plaza-dev/uru-databases-2-go-service-common/utils"
	"strings"
)

type UnexpectedSigningMethodError struct {
	Algorithm interface{}
}

// Error returns a formatted error message for UnexpectedSigningMethodError
func (u UnexpectedSigningMethodError) Error() (message string) {
	formattedAlgorithm := utils.AddBrackets(u.Algorithm.(string))
	return strings.Join([]string{"Unexpected signing method", formattedAlgorithm}, " ")
}
