package validator

import (
	"errors"
)

var (
	InvalidTokenError            = errors.New("invalid token")
	UnexpectedSigningMethodError = errors.New("unexpected signing method")
	InvalidClaimsError           = errors.New("invalid claims")
	IRTNotValidError             = errors.New("irt not valid")
	IdentifierNotValidError      = errors.New("jwt_id not valid")
	MustBeAccessTokenError       = errors.New("must be access token")
	MustBeRefreshTokenError      = errors.New("must be refresh token")
)
