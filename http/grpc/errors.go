package grpc

import (
	"errors"
)

var (
	InternalServerError                   = errors.New("internal server error")
	AuthorizationMetadataInvalidError     = errors.New("authorization metadata invalid")
	AuthorizationMetadataNotProvidedError = errors.New("authorization metadata is not provided")
)
