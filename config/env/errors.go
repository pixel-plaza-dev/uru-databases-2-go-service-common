package env

import "errors"

var (
	FailedToLoadEnvironmentVariablesError = errors.New("failed to load environment variables")
)
