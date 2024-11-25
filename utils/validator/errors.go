package validator

import "errors"

var (
	FieldNotFoundError      = errors.New("field not found")
	InvalidBirthDateError   = errors.New("invalid birth date")
	InvalidMailAddressError = errors.New("invalid mail address")
	EmptyStringFieldError   = errors.New("string field is empty")
)
