package validator

import "errors"

var (
	NilValidationsError     = errors.New("nil validations")
	NilFieldsError          = errors.New("nil fields")
	NilDataError            = errors.New("nil data")
	FieldNotFoundError      = errors.New("field not found")
	InvalidBirthdateError   = errors.New("invalid birthdate")
	InvalidMailAddressError = errors.New("invalid mail address")
	EmptyStringFieldError   = errors.New("string field is empty")
)
