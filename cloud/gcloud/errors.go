package gcloud

import "errors"

var (
	NilGoogleCredentialsError          = errors.New("google credentials cannot be nil")
	NilTokenSourceError                = errors.New("token source cannot be nil")
	FailedToLoadGoogleCredentialsError = errors.New("failed to load google credentials")
	FailedToCreateTokenSourceError     = errors.New("failed to create token source")
)
