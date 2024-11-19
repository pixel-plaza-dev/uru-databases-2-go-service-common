package gcloud

import "errors"

var (
	FailedToGenerateDefaultGoogleCredentialsError = errors.New("failed to create default google credentials")
	FailedToCreateTokenSourceError                = errors.New("failed to create token source")
)
