package grpc

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	InternalServerError                   = status.Error(codes.Internal, "internal server error")
	ServiceUnavailable                    = errors.New("service unavailable")
	Unauthenticated                       = errors.New("unauthenticated")
	Unauthorized                          = errors.New("unauthorized")
	AuthorizationMetadataInvalidError     = errors.New("authorization metadata invalid")
	AuthorizationMetadataNotProvidedError = status.Error(
		codes.Unauthenticated, "authorization metadata is not provided",
	)
)
