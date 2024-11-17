package grpc

import (
	enverror "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/config/env/error"
	"os"
)

// LoadServiceURI load service URI from environment variables
func LoadServiceURI(key string) (uri string, err error) {
	// Get environment variable
	uri, uriExists := os.LookupEnv(key)
	if !uriExists {
		return "", enverror.VariableNotFoundError{Key: key}
	}
	return uri, nil
}
