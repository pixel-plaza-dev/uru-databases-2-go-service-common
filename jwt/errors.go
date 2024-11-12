package jwt

import "errors"

var (
	UnableToReadKeyFileError     = errors.New("unable to read private key file")
	UnableToParsePrivateKeyError = errors.New("unable to parse private key")
	UnableToParsePublicKeyError  = errors.New("unable to parse public key")
	InvalidTokenError            = errors.New("invalid token")
	UnableToIssueTokenError      = errors.New("unable to issue token")
	UnexpectedSigningMethodError = errors.New("unexpected signing method")
	InvalidClaimsError           = errors.New("invalid claims")
	TokenExpiredError            = errors.New("token expired")
)
