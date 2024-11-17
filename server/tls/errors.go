package tls

import "errors"

var (
	FailedToAddCAPemError              = errors.New("failed to add server ca's certificate")
	FailedToLoadSystemCredentialsError = errors.New("failed to load system credentials")
)
