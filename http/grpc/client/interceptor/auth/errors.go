package auth

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	MissingTokenError        = status.Error(codes.Unauthenticated, "missing token")
	UnexpectedTokenTypeError = status.Error(codes.Unauthenticated, "unexpected token type")
)
