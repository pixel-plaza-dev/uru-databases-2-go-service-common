package grpc

import "errors"

var (
	InternalServerError                = errors.New("internal server error")
	ServiceUnavailable                 = errors.New("service unavailable")
	Unauthenticated                    = errors.New("unauthenticated")
	Unauthorized                       = errors.New("unauthorized")
	FailedToAddCAPemError              = errors.New("failed to add server ca's certificate")
	FailedToCreateTokenSourceError     = errors.New("failed to create token source")
	FailedToLoadSystemCredentialsError = errors.New("failed to load system credentials")
)
