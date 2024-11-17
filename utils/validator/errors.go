package validator

import "errors"

var (
	FieldNotFoundError      = errors.New("field not found")
	InvalidMailAddressError = errors.New("invalid mail address")
	EmptyStringFieldError   = errors.New("string field is empty")
)
