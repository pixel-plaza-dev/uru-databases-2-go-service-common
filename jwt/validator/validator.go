package validator

import (
	"crypto"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	commonjwt "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/jwt"
	"os"
)

// Validator does parsing and validation of JWT token
type (
	Validator interface {
		GetToken(tokenString string) (*jwt.Token, error)
		GetClaims(tokenString string) (*jwt.MapClaims, error)
	}

	DefaultValidator struct {
		key            *crypto.PublicKey
		validateClaims func(*jwt.MapClaims) (*jwt.MapClaims, error)
	}
)

// NewDefaultValidator returns a new validator by parsing the given file path as an ED25519 public key
func NewDefaultValidator(publicKeyPath string, validateClaims func(*jwt.MapClaims) (*jwt.MapClaims, error)) (*DefaultValidator, error) {
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

	return &DefaultValidator{
		key:            &key,
		validateClaims: validateClaims,
	}, nil
}

// GetToken parses the given JWT token string
func (d *DefaultValidator) GetToken(tokenString string) (*jwt.Token, error) {
	// Parse JWT and verify signature
	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {
			// Check to see if the token uses the expected signing method
			if _, ok := token.Method.(*jwt.SigningMethodEd25519); !ok {
				return nil, commonjwt.UnexpectedSigningMethodError
			}
			return d.key, nil
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

	// Get the claims from the token
	return token, nil
}

// GetClaims parses and validates the given JWT token string
func (d *DefaultValidator) GetClaims(tokenString string) (*jwt.MapClaims, error) {
	// Get the token
	token, err := d.GetToken(tokenString)
	if err != nil {
		return nil, err
	}

	// Get token claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, commonjwt.InvalidClaimsError
	}

	return d.validateClaims(&claims)
}
