package validator

import (
	"net/mail"
)

// ValidMailAddress checks if the mail address is valid
func ValidMailAddress(address string) (string, error) {
	addr, err := mail.ParseAddress(address)
	if err != nil {
		return "", InvalidMailAddressError
	}

	return addr.Address, nil
}
