package grpc

type AuthorizationIdx int

const (
	TokenIdx AuthorizationIdx = iota
)

// Int returns the integer value of the index
func (i AuthorizationIdx) Int() int {
	return int(i)
}

const (
	// AuthorizationMetadataKey is the key for the authorization metadata
	AuthorizationMetadataKey = "authorization"

	// BearerPrefix is the prefix for the bearer token
	BearerPrefix = "Bearer"

	// JwtCtxClaimsKey is the key for the JWT context claims
	JwtCtxClaimsKey = "jwt_claims"
)
