package env

import (
	enverror "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/config/env/error"
	"os"
)

// LoadVariable load variable from environment variables
func LoadVariable(key string) (uri string, err error) {
	// Get environment variable
	variable, variableExists := os.LookupEnv(key)
	if !variableExists {
		return "", enverror.VariableNotFoundError{Key: key}
	}
	return variable, nil
}
