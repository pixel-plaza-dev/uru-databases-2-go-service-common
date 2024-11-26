package context

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pixel-plaza-dev/uru-databases-2-go-service-common/http/grpc"
)

// SetCtxTokenString sets the token string to the context
func SetCtxTokenString(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, grpc.AuthorizationMetadataKey, token)
}

// SetCtxTokenClaims sets the token claims to the context
func SetCtxTokenClaims(ctx context.Context, claims jwt.Claims) context.Context {
	return context.WithValue(ctx, grpc.CtxTokenClaimsKey, claims)
}

// GetCtxTokenString gets the token string from the context
func GetCtxTokenString(ctx context.Context) (string, error) {
	// Get the token from the context
	value := ctx.Value(grpc.AuthorizationMetadataKey)
	if value == nil {
		return "", MissingTokenInContextError
	}

	// Check the type of the value
	token, ok := value.(string)
	if !ok {
		return "", UnexpectedTokenTypeInContextError
	}

	return token, nil
}

// GetCtxTokenClaims gets the token claims from the context
func GetCtxTokenClaims(ctx context.Context) (jwt.Claims, error) {
	// Get the claims from the context
	value := ctx.Value(grpc.CtxTokenClaimsKey)
	if value == nil {
		return nil, MissingTokenClaimsInContextError
	}

	// Check the type of the value
	claims, ok := value.(jwt.Claims)
	if !ok {
		return nil, UnexpectedTokenClaimsTypeInContextError
	}
	return claims, nil
}

// GetCtxTokenClaimsSubject gets the token claims subject from the context
func GetCtxTokenClaimsSubject(ctx context.Context) (string, error) {
	// Get the claims from the context
	claims, err := GetCtxTokenClaims(ctx)
	if err != nil {
		return "", err
	}

	// Get the subject from the claims
	subject, err := claims.GetSubject()
	if err != nil {
		return "", MissingTokenClaimsSubjectInContextError
	}
	return subject, nil
}
