package validator

type InvalidMailAddressError struct{}

// Error returns a formatted error message for InvalidMailAddressError
func (i InvalidMailAddressError) Error() (message string) {
	return "Mail address is invalid"
}
