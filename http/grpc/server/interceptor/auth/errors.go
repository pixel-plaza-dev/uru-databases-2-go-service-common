package auth

import (
	"errors"
)

var (
	MetadataNotProvidedError  = errors.New("metadata is not provided")
	GRPCInterceptionsNilError = errors.New("nil grpc interceptions map")
)
