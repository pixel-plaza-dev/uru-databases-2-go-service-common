package auth

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	// MissingTokenError is the error for missing JWT token
	MissingTokenError = status.Error(codes.Unauthenticated, "missing token")

	// UnexpectedTokenTypeError is the error for unexpected token type
	UnexpectedTokenTypeError = status.Error(codes.Unauthenticated, "unexpected token type")
)
