package issuer

import (
	"crypto"
	"github.com/golang-jwt/jwt/v5"
	commonjwt "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/crypto/jwt"
	"time"
)

type (
	// Issuer is the interface for JWT issuing
	Issuer interface {
		IssueToken(claims *jwt.MapClaims) (string, error)
	}

	// DefaultIssuer handles JWT issuing
	DefaultIssuer struct {
		key *crypto.PrivateKey
	}
)

// NewDefaultIssuer creates a new issuer by parsing the given path as an ED25519 private key
func NewDefaultIssuer(privateKey []byte) (*DefaultIssuer, error) {
	// Parse the private key
	key, err := jwt.ParseEdPrivateKeyFromPEM(privateKey)
	if err != nil {
		return nil, commonjwt.UnableToParsePrivateKeyError
	}

	return &DefaultIssuer{
		key: &key,
	}, nil
}

// GenerateClaims generates a new claims object
func (i *DefaultIssuer) GenerateClaims(
	jwtId string, userId string, expirationTime time.Time,
) *jwt.MapClaims {
	return &jwt.MapClaims{
		"exp": expirationTime.Unix(),
		"iat": time.Now().Unix(),
		"jti": jwtId,
		"sub": userId,
	}
}

// IssueToken issues a new token for the given user with the given roles
func (i *DefaultIssuer) IssueToken(claims *jwt.MapClaims) (string, error) {
	// Create a new token with the claims
	token := jwt.NewWithClaims(&jwt.SigningMethodEd25519{}, claims)

	// Sign and get the complete encoded token as a string using the private key
	tokenString, err := token.SignedString(i.key)
	if err != nil {
		return "", commonjwt.UnableToIssueTokenError
	}

	return tokenString, nil
}
