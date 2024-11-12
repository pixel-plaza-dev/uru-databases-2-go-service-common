package jwt

import (
	enverror "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/env/error"
	"os"
)

// LoadJwtKey loads the JWT key from the environment
func LoadJwtKey(jwtKey string) (string, error) {
	// Get JWT key from environment
	jwt, exists := os.LookupEnv(jwtKey)
	if !exists {
		return "", enverror.VariableNotFoundError{Key: jwtKey}
	}

	return jwt, nil
}

// ReadJwtKey reads the JWT key from the given path
func ReadJwtKey(path string) ([]byte, error) {
	jwt, err := os.ReadFile(path)
	if err != nil {
		return nil, UnableToReadKeyFileError
	}
	return jwt, nil
}
