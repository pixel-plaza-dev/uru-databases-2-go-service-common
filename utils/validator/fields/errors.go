package fields

import (
	"errors"
)

var (
	MissingProtobufTagError         = "missing protobuf tag: %s"
	MissingProtobufNameError        = "missing protobuf name: %s"
	NilStructFieldsValidationsError = errors.New("nil struct fields validations")
	NilDataError                    = errors.New("nil data")
	NilStructFieldsToValidateError  = errors.New("nil struct fields to validate")
	FieldNotFoundError              = errors.New("field not found")
	InvalidBirthdateError           = errors.New("invalid birthdate")
	InvalidMailAddressError         = errors.New("invalid mail address")
	EmptyStringFieldError           = errors.New("string field is empty")
)
