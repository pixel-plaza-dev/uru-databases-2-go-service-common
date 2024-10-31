package grpc

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
)

// SetContextClaims sets the claims in the context
func SetContextClaims(ctx *context.Context, claims jwt.Claims) context.Context {
	return context.WithValue(*ctx, JwtCtxClaimsKey, claims)
}
