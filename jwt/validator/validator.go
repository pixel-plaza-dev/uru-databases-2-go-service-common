package validator

import (
	"crypto"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	commonjwt "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/jwt"
	"os"
)

// Validator does parsing and validation of JWT token
type Validator struct {
	key            *crypto.PublicKey
	validateClaims func(*jwt.Token) (*jwt.Token, error)
}

// NewValidator returns a new validator by parsing the given file path as an ED25519 public key
func NewValidator(publicKeyPath string, validateClaims func(*jwt.Token) (*jwt.Token, error)) (*Validator, error) {
	// Read the public key file
	keyBytes, err := os.ReadFile(publicKeyPath)
	if err != nil {
		return nil, commonjwt.UnableToReadPublicKeyFileError
	}

	// Parse the public key
	key, err := jwt.ParseEdPublicKeyFromPEM(keyBytes)
	if err != nil {
		return nil, commonjwt.UnableToParsePublicKeyError
	}

	return &Validator{
		key:            &key,
		validateClaims: validateClaims,
	}, nil
}

// GetToken attempts to get a token from the given string
func (v *Validator) GetToken(tokenString string) (*jwt.Token, error) {
	// Parse JWT and verify signature
	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {
			// Check to see if the token uses the expected signing method
			if _, ok := token.Method.(*jwt.SigningMethodEd25519); !ok {
				return nil, commonjwt.UnexpectedSigningMethodError
			}
			return v.key, nil
		})
	if err != nil {
		switch {
		case errors.Is(err, commonjwt.UnexpectedSigningMethodError):
		case errors.Is(err, jwt.ErrSignatureInvalid):
		case errors.Is(err, jwt.ErrTokenExpired):
		case errors.Is(err, jwt.ErrTokenMalformed):
			return nil, err
		default:
			return nil, commonjwt.InvalidTokenError
		}
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, commonjwt.InvalidTokenError
	}

	// Validate the token claims with the given function
	return v.validateClaims(token)
}
