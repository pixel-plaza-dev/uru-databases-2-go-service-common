package validator

import (
	"crypto"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	commonjwt "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/crypto/jwt"
)

// Validator does parsing and validation of JWT token
type (
	Validator interface {
		GetToken(tokenString string) (*jwt.Token, error)
		GetClaims(tokenString string) (*jwt.MapClaims, error)
		GetValidatedClaims(token string, mustBeRefreshToken bool) (*jwt.MapClaims, error)
	}

	DefaultValidator struct {
		key          *crypto.PublicKey
		isTokenValid func(token string, isRefreshToken bool) bool
	}
)

// NewDefaultValidator returns a new validator by parsing the given file path as an ED25519 public key
func NewDefaultValidator(
	publicKey []byte, isTokenValid func(token string, isRefreshToken bool) bool,
) (*DefaultValidator, error) {
	// Parse the public key
	key, err := jwt.ParseEdPublicKeyFromPEM(publicKey)
	if err != nil {
		return nil, commonjwt.UnableToParsePublicKeyError
	}

	return &DefaultValidator{
		key:          &key,
		isTokenValid: isTokenValid,
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
		},
	)
	if err != nil {
		switch {
		case errors.Is(err, commonjwt.UnexpectedSigningMethodError):
		case errors.Is(err, jwt.ErrSignatureInvalid):
		case errors.Is(err, jwt.ErrTokenExpired):
		case errors.Is(err, jwt.ErrTokenNotValidYet):
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
func (d *DefaultValidator) GetClaims(tokenString string) (
	*jwt.MapClaims, error,
) {
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

	return &claims, nil
}

// ValidateClaims validates the given claims
func (d *DefaultValidator) ValidateClaims(token string, claims *jwt.MapClaims, mustBeRefreshToken bool) (*jwt.MapClaims, error) {
	// Check if is a refresh token
	irt, ok := (*claims)[commonjwt.IsRefreshTokenClaim].(bool)
	if !ok {
		return nil, commonjwt.IRTNotValidError
	}

	// Check if it must be a refresh token
	if irt != mustBeRefreshToken {
		if mustBeRefreshToken {
			return nil, commonjwt.MustBeRefreshTokenError
		}
		return nil, commonjwt.MustBeAccessTokenError
	}

	// Check if the token is valid
	if !d.isTokenValid(token, irt) {
		return nil, commonjwt.InvalidTokenError
	}

	return claims, nil
}

// GetValidatedClaims parses, validates and returns the claims of the given JWT token string
func (d *DefaultValidator) GetValidatedClaims(token string, mustBeRefreshToken bool) (
	*jwt.MapClaims, error,
) {
	// Get the claims
	claims, err := d.GetClaims(token)
	if err != nil {
		return nil, err
	}

	// Validate the claims
	return d.ValidateClaims(token, claims, mustBeRefreshToken)
}
