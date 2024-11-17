package grpc

import "errors"

var (
	InternalServerError               = errors.New("internal server error")
	ServiceUnavailable                = errors.New("service unavailable")
	Unauthenticated                   = errors.New("unauthenticated")
	Unauthorized                      = errors.New("unauthorized")
	AuthorizationMetadataInvalidError = errors.New("authorization metadata invalid")
	MissingBearerPrefixError          = errors.New("missing bearer prefix")
	MissingBearerTokenError           = errors.New("missing bearer token")
)