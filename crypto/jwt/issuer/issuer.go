package issuer

import (
	"crypto"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
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

// GetExpirationTime returns the expiration time for the given duration
func (i *DefaultIssuer) GetExpirationTime(
	issuedTime time.Time,
	duration time.Duration,
) time.Time {
	return issuedTime.Add(duration)
}

// GenerateClaims generates a new claims object
func (i *DefaultIssuer) GenerateClaims(
	jwtId string,
	userId string,
	userUUID uuid.UUID,
	issuedTime time.Time,
	expirationTime time.Time,
	isRefreshToken bool,
) *jwt.MapClaims {
	return &jwt.MapClaims{
		"exp":                         expirationTime.Unix(),
		"iat":                         issuedTime.Unix(),
		commonjwt.IdClaim:             jwtId,
		commonjwt.UserIdClaim:         userId,
		commonjwt.UserSharedIdClaim:   userUUID.String(),
		commonjwt.IsRefreshTokenClaim: isRefreshToken,
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
