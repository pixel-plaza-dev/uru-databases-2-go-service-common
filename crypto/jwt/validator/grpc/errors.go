package grpc

import (
	"errors"
)

var (
	NilTokenValidatorError = errors.New("token validator cannot be nil")
)
