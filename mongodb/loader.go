package mongodb

import (
	enverror "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/env/error"
	"os"
)

// LoadMongoDBURI load MongoDB URI from environment variables
func LoadMongoDBURI(key string) (uri string, err error) {
	// Get environment variable
	uri, uriExists := os.LookupEnv(key)
	if !uriExists {
		return "", enverror.VariableNotFoundError{Key: key}
	}
	return uri, nil
}

// LoadMongoDBName load MongoDB database name from environment variables
func LoadMongoDBName(key string) (database string, err error) {
	// Get environment variable
	database, databaseExists := os.LookupEnv(key)
	if !databaseExists {
		return "", enverror.VariableNotFoundError{Key: key}
	}

	return database, nil
}
