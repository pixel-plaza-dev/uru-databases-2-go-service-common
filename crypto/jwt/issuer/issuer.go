package issuer

import (
	"github.com/golang-jwt/jwt/v5"
)

// Issuer is the interface for JWT issuing
type Issuer interface {
	IssueToken(claims *jwt.MapClaims) (string, error)
}
