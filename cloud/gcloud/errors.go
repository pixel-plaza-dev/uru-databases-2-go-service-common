package gcloud

import "errors"

var (
	FailedToLoadGoogleCredentialsError = errors.New("failed to load google credentials")
	FailedToCreateTokenSourceError     = errors.New("failed to create token source")
)
