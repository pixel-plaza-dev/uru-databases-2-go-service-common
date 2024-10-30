package validator

import "errors"

var FieldNotFoundError = errors.New("field not found")
var InvalidMailAddressError = errors.New("invalid mail address")
var EmptyStringFieldError = errors.New("string field is empty")
