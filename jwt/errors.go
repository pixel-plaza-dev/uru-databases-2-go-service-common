package jwt

import "errors"

var UnableToReadPrivateKeyFileError = errors.New("unable to read private key file")
var UnableToReadPublicKeyFileError = errors.New("unable to read public key file")
var UnableToParsePrivateKeyError = errors.New("unable to parse private key")
var UnableToParsePublicKeyError = errors.New("unable to parse public key")
var InvalidTokenError = errors.New("invalid token")
var UnableToIssueTokenError = errors.New("unable to issue token")
var UnexpectedSigningMethodError = errors.New("unexpected signing method")
