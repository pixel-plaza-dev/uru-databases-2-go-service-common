package validator

import "errors"

var (
	FieldNotFoundError      = errors.New("field not found")
	InvalidBirthdateError   = errors.New("invalid birthdate")
	InvalidMailAddressError = errors.New("invalid mail address")
	EmptyStringFieldError   = errors.New("string field is empty")
)
