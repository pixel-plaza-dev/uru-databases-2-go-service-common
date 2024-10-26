package validator

import (
	"github.com/pixel-plaza-dev/uru-databases-2-go-service-common/custom_error_response/validator"
	"net/mail"
)

// ValidMailAddress checks if the mail address is valid
func ValidMailAddress(address string) (string, error) {
	addr, err := mail.ParseAddress(address)
	if err != nil {
		return "", validator.InvalidMailAddressError{}
	}

	return addr.Address, nil
}
