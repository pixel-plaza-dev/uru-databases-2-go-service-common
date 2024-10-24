package validator

import "net/mail"

// ValidMailAddress checks if the mail address is valid
func ValidMailAddress(address string) (string, bool) {
	addr, err := mail.ParseAddress(address)
	if err != nil {
		return "", false
	}
	return addr.Address, true
}
