package context

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pixel-plaza-dev/uru-databases-2-go-service-common/server/grpc"
)

// SetCtxClaims sets the claims in the context
func SetCtxClaims(ctx *context.Context, claims jwt.Claims) context.Context {
	return context.WithValue(*ctx, grpc.JwtCtxClaimsKey, claims)
}
