package jwt

import (
	"crypto"
	"github.com/golang-jwt/jwt/v5"
	customJwtErrror "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/custom_error/jwt"
	"github.com/pixel-plaza-dev/uru-databases-2-go-service-common/logger"
	"os"
)

// Issuer handles JWT issuing
type Issuer struct {
	name   string
	key    *crypto.PrivateKey
	logger *logger.JwtIssuerLogger
}

// NewIssuer creates a new issuer by parsing the given path as an ED25519 private key
func NewIssuer(name string, privateKeyPath string) (*Issuer, error) {
	// Create the issuer logger
	issuerLogger := logger.NewJwtIssuerLogger(name)

	// Read the private key file
	keyBytes, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return nil, customJwtErrror.UnableToParseKeyError{Err: err, KeyType: customJwtErrror.PrivateKey}
	}

	// Parse the private key
	key, err := jwt.ParseEdPrivateKeyFromPEM(keyBytes)
	if err != nil {
		return nil, customJwtErrror.UnableToParseKeyError{Err: err, KeyType: customJwtErrror.PrivateKey}
	}

	return &Issuer{
		key:    &key,
		name:   name,
		logger: issuerLogger,
	}, nil
}

// IssueToken issues a new token for the given user with the given roles
func (i *Issuer) IssueToken(claims *jwt.MapClaims) (string, error) {
	// Create a new token with the claims
	token := jwt.NewWithClaims(&jwt.SigningMethodEd25519{}, claims)

	// Sign and get the complete encoded token as a string using the private key
	tokenString, err := token.SignedString(i.key)
	if err != nil {
		return "", customJwtErrror.UnableToIssueTokenError{Err: err}
	}

	return tokenString, nil
}
