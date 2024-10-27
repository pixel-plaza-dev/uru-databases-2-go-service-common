package validator

import (
	"github.com/pixel-plaza-dev/uru-databases-2-go-service-common/validator/error/response"
	"net/mail"
)

// ValidMailAddress checks if the mail address is valid
func ValidMailAddress(address string) (string, error) {
	addr, err := mail.ParseAddress(address)
	if err != nil {
		return "", response.InvalidMailAddressError{}
	}

	return addr.Address, nil
}
