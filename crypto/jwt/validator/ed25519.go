package validator

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	commonflag "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/config/flag"
	commonjwt "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/crypto/jwt"
	commonjwtvalidatorgrpc "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/crypto/jwt/validator/grpc"
	pbtypesgrpc "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/grpc"
	"golang.org/x/crypto/ed25519"
)

// Ed25519Validator struct
type Ed25519Validator struct {
	ed25519Key     *ed25519.PublicKey
	tokenValidator commonjwtvalidatorgrpc.TokenValidator
	mode           *commonflag.ModeFlag
}

// NewEd25519Validator returns a new validator by parsing the given file path as an ED25519 public key
func NewEd25519Validator(
	publicKey []byte, tokenValidator commonjwtvalidatorgrpc.TokenValidator, mode *commonflag.ModeFlag,
) (*Ed25519Validator, error) {
	// Check if either the token validator or the mode flag is nil
	if tokenValidator == nil {
		return nil, commonjwtvalidatorgrpc.NilTokenValidatorError
	}
	if mode == nil {
		return nil, commonflag.NilModeFlagError
	}

	// Parse the public key
	key, err := jwt.ParseEdPublicKeyFromPEM(publicKey)
	if err != nil {
		return nil, commonjwt.UnableToParsePublicKeyError
	}

	// Ensure the key is of type ED25519 public key
	ed25519Key, ok := key.(ed25519.PublicKey)
	if !ok {
		return nil, commonjwt.InvalidKeyTypeError
	}

	return &Ed25519Validator{
		ed25519Key:     &ed25519Key,
		tokenValidator: tokenValidator,
		mode:           mode,
	}, nil
}

// GetToken parses the given JWT token string
func (d *Ed25519Validator) GetToken(tokenString string) (*jwt.Token, error) {
	// Parse JWT and verify signature
	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {
			// Check to see if the token uses the expected signing method
			if _, ok := token.Method.(*jwt.SigningMethodEd25519); !ok {
				return nil, UnexpectedSigningMethodError
			}
			return *d.ed25519Key, nil
		},
	)
	if err != nil {
		if d.mode.IsDev() {
			return nil, err
		}

		switch {
		case errors.Is(err, UnexpectedSigningMethodError):
		case errors.Is(err, jwt.ErrSignatureInvalid):
		case errors.Is(err, jwt.ErrTokenExpired):
		case errors.Is(err, jwt.ErrTokenNotValidYet):
		case errors.Is(err, jwt.ErrTokenMalformed):
			return nil, err
		default:
			return nil, InvalidTokenError
		}
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, InvalidTokenError
	}

	// Get the claims from the token
	return token, nil
}

// GetClaims parses and validates the given JWT token string
func (d *Ed25519Validator) GetClaims(tokenString string) (
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
		return nil, InvalidClaimsError
	}

	return &claims, nil
}

// ValidateClaims validates the given claims
func (d *Ed25519Validator) ValidateClaims(
	token string,
	claims *jwt.MapClaims,
	interception pbtypesgrpc.Interception,
) (*jwt.MapClaims, error) {
	// Check if the claims are nil
	if claims == nil {
		return nil, NilJwtClaimsError
	}

	// Check if is a refresh token
	irt, ok := (*claims)[commonjwt.IsRefreshTokenClaim].(bool)
	if !ok {
		return nil, IRTNotValidError
	}

	// Get the JWT Identifier
	jwtId, ok := (*claims)[commonjwt.IdClaim].(string)
	if !ok {
		return nil, IdentifierNotValidError
	}

	// Check if it must be a refresh token
	if !irt && interception == pbtypesgrpc.RefreshToken {
		return nil, MustBeRefreshTokenError
	}

	// Check if it must be an access token
	if irt && interception == pbtypesgrpc.AccessToken {
		return nil, MustBeAccessTokenError
	}

	// Check if the token is valid
	isValid, err := d.tokenValidator.IsTokenValid(token, jwtId, irt)
	if err != nil {
		return nil, err
	}
	if !isValid {
		return nil, InvalidTokenError
	}

	return claims, nil
}

// GetValidatedClaims parses, validates and returns the claims of the given JWT token string
func (d *Ed25519Validator) GetValidatedClaims(
	token string,
	interception pbtypesgrpc.Interception,
) (
	*jwt.MapClaims, error,
) {
	// Get the claims
	claims, err := d.GetClaims(token)
	if err != nil {
		return nil, err
	}

	// Validate the claims
	return d.ValidateClaims(token, claims, interception)
}
