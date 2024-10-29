package issuer

import (
	"crypto"
	"github.com/golang-jwt/jwt/v5"
	jwterror "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/jwt/error"
	"os"
)

// Issuer handles JWT issuing
type Issuer struct {
	key *crypto.PrivateKey
}

// NewIssuer creates a new issuer by parsing the given path as an ED25519 private key
func NewIssuer(privateKeyPath string) (*Issuer, error) {
	// Read the private key file
	keyBytes, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return nil, jwterror.UnableToParseKeyError{Err: err, KeyType: jwterror.PrivateKey}
	}

	// Parse the private key
	key, err := jwt.ParseEdPrivateKeyFromPEM(keyBytes)
	if err != nil {
		return nil, jwterror.UnableToParseKeyError{Err: err, KeyType: jwterror.PrivateKey}
	}

	return &Issuer{
		key: &key,
	}, nil
}

// IssueToken issues a new token for the given user with the given roles
func (i *Issuer) IssueToken(claims *jwt.MapClaims) (string, error) {
	// Create a new token with the claims
	token := jwt.NewWithClaims(&jwt.SigningMethodEd25519{}, claims)

	// Sign and get the complete encoded token as a string using the private key
	tokenString, err := token.SignedString(i.key)
	if err != nil {
		return "", jwterror.UnableToIssueTokenError{Err: err}
	}

	return tokenString, nil
}
