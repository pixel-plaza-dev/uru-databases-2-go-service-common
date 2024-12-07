package issuer

import (
	"github.com/golang-jwt/jwt/v5"
)

type (
	// Issuer is the interface for JWT issuing
	Issuer interface {
		IssueToken(claims *jwt.MapClaims) (string, error)
	}
)
