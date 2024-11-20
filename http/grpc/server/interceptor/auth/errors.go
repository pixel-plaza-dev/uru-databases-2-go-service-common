package auth

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	MetadataNotProvidedError = status.Error(
		codes.Unauthenticated, "metadata is not provided",
	)
	GRPCInterceptionsNilError = errors.New("nil gRPC interceptions map")
)
