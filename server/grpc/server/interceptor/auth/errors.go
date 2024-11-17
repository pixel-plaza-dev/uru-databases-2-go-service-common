package auth

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	MetadataNotProvidedError              = status.Error(codes.Unauthenticated, "metadata is not provided")
	AuthorizationMetadataNotProvidedError = status.Error(codes.Unauthenticated, "authorization metadata is not provided")
)
