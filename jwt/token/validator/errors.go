package validator

import (
	"errors"
)

var (
	NilClaimsValidator               = errors.New("claims validator cannot be nil")
	IsRefreshTokenClaimNotValidError = errors.New("irt not valid")
	IdClaimNotValidError             = errors.New("jwt_id not valid")
	MustBeAccessTokenError           = errors.New("must be access token")
	MustBeRefreshTokenError          = errors.New("must be refresh token")
	TokenExpiredError                = errors.New("token expired")
)
