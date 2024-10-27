package validator

import (
	"crypto"
	"github.com/golang-jwt/jwt/v5"
	jwterror "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/jwt/error"
	"os"
)

// Validator does parsing and validation of JWT token
type Validator struct {
	name   string
	key    *crypto.PublicKey
	logger *Logger
}

// NewValidator returns a new validator by parsing the given file path as an ED25519 public key
func NewValidator(name string, publicKeyPath string) (*Validator, error) {
	// Create the validator logger
	validatorLogger := NewJwtValidatorLogger(name)

	// Read the public key file
	keyBytes, err := os.ReadFile(publicKeyPath)
	if err != nil {
		return nil, jwterror.UnableToParseKeyError{Err: err, KeyType: jwterror.PublicKey}
	}

	// Parse the public key
	key, err := jwt.ParseEdPublicKeyFromPEM(keyBytes)
	if err != nil {
		return nil, jwterror.UnableToParseKeyError{Err: err, KeyType: jwterror.PublicKey}
	}

	return &Validator{
		key:    &key,
		name:   name,
		logger: validatorLogger,
	}, nil
}

// GetToken attempts to get a token from the given string
func (v *Validator) GetToken(tokenString string, validateClaims func(*jwt.Token) (*jwt.Token, error)) (*jwt.Token, error) {
	// Parse JWT and verify signature
	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {
			// Check to see if the token uses the expected signing method
			if _, ok := token.Method.(*jwt.SigningMethodEd25519); !ok {
				alg := token.Header["alg"]
				return nil, jwterror.UnexpectedSigningMethodError{Algorithm: alg}
			}
			return v.key, nil
		})
	if err != nil {
		return nil, jwterror.UnableToParseTokenError{Err: err}
	}

	// Validate the token claims with the given function
	return validateClaims(token)
}
