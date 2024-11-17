package auth

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	MetadataNotProvidedError            = status.Error(codes.Unauthenticated, "metadata is not provided")
	AuthorizationHeaderNotProvidedError = status.Error(codes.Unauthenticated, "authorization header is not provided")
)
