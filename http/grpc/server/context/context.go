package context

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	commonjwt "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/crypto/jwt"
	"github.com/pixel-plaza-dev/uru-databases-2-go-service-common/http/grpc"
	"google.golang.org/grpc/peer"
	"net"
	"strings"
)

// SetCtxTokenString sets the token string to the context
func SetCtxTokenString(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, grpc.AuthorizationMetadataKey, token)
}

// SetCtxTokenClaims sets the token claims to the context
func SetCtxTokenClaims(
	ctx context.Context,
	claims *jwt.MapClaims,
) context.Context {
	return context.WithValue(ctx, grpc.CtxTokenClaimsKey, claims)
}

// GetCtxTokenString gets the token string from the context
func GetCtxTokenString(ctx context.Context) (string, error) {
	// Get the token from the context
	value := ctx.Value(grpc.AuthorizationMetadataKey)
	if value == nil {
		return "", MissingTokenError
	}

	// Check the type of the value
	token, ok := value.(string)
	if !ok {
		return "", UnexpectedTokenTypeError
	}

	return token, nil
}

// GetCtxTokenClaims gets the token claims from the context
func GetCtxTokenClaims(ctx context.Context) (*jwt.MapClaims, error) {
	// Get the claims from the context
	value := ctx.Value(grpc.CtxTokenClaimsKey)
	if value == nil {
		return nil, MissingTokenClaimsError
	}

	// Check the type of the value
	claims, ok := value.(*jwt.MapClaims)
	if !ok {
		return nil, UnexpectedTokenClaimsTypeError
	}
	return claims, nil
}

// GetCtxTokenClaimsUserId gets the token claims user ID from the context
func GetCtxTokenClaimsUserId(ctx context.Context) (string, error) {
	// Get the claims from the context
	claims, err := GetCtxTokenClaims(ctx)
	if err != nil {
		return "", err
	}

	// Get the user ID from the claims
	userId, ok := (*claims)[commonjwt.UserIdClaim].(string)
	if !ok {
		return "", MissingTokenClaimsUserIdError
	}
	return userId, nil
}

// GetCtxTokenClaimsJwtId gets the token claims JWT ID from the context
func GetCtxTokenClaimsJwtId(ctx context.Context) (string, error) {
	// Get the claims from the context
	claims, err := GetCtxTokenClaims(ctx)
	if err != nil {
		return "", err
	}

	// Get the JWT ID from the claims
	jwtId, ok := (*claims)[commonjwt.IdClaim].(string)
	if !ok {
		return "", MissingTokenClaimsIdError
	}
	return jwtId, nil
}

// GetClientIP extracts the client IP address from the context
func GetClientIP(ctx context.Context) (string, error) {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return "", FailedToGetPeerFromContextError
	}

	// Get the IP address from the peer address
	addr := p.Addr.String()
	ip, _, err := net.SplitHostPort(addr)
	if err != nil {
		return "", err
	}

	// Remove any surrounding brackets from IPv6 addresses
	ip = strings.Trim(ip, "[]")

	return ip, nil
}
