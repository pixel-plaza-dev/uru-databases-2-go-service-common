package grpc

import "errors"

var (
	InternalServerError   = errors.New("internal server error")
	ServiceUnavailable    = errors.New("service unavailable")
	Unauthenticated       = errors.New("unauthenticated")
	Unauthorized          = errors.New("unauthorized")
	FailedToAddCAPemError = errors.New("failed to add server CA's certificate")
)
