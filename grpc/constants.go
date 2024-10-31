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
	// AuthorizationHeaderKey is the key for the authorization header
	AuthorizationHeaderKey = "authorization"

	// JwtCtxClaimsKey is the key for the JWT context claims
	JwtCtxClaimsKey = "jwt_claims"
)
