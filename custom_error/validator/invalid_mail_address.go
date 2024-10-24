package validator

import (
	"github.com/pixel-plaza-dev/uru-databases-2-go-service-common/utils"
	"strings"
)

type InvalidMailAddressError struct {
	MailAddress string
}

// Error returns a formatted error message for InvalidMailAddressError
func (i InvalidMailAddressError) Error() (message string) {
	formattedMailAddress := utils.AddBrackets(i.MailAddress)
	return strings.Join([]string{"Mail address is invalid", formattedMailAddress}, " ")
}
