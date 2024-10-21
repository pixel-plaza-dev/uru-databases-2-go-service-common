package loader

import (
	"github.com/pixel-plaza-dev/uru-databases-2-go-service-common/logger"
	"os"
	"strings"
)

// LoadServicePort load Service port from environment variables
func LoadServicePort(key string, logger *logger.EnvironmentLogger) (port string, formattedPort string) {
	// Get environment variable
	port, exists := os.LookupEnv(key)
	if !exists {
		logger.VariableNotFound(key)
	}

	// Build port string
	var portBuilder strings.Builder
	portBuilder.WriteString(":")
	portBuilder.WriteString(port)

	return port, portBuilder.String()
}

// LoadMongoDBURI load MongoDB URI from environment variables
func LoadMongoDBURI(key string, logger *logger.EnvironmentLogger) (uri string) {
	// Get environment variable
	uri, uriExists := os.LookupEnv(key)
	if !uriExists {
		logger.VariableNotFound(key)
	}
	return uri
}

// LoadMongoDBName load MongoDB database name from environment variables
func LoadMongoDBName(key string, logger *logger.EnvironmentLogger) (database string) {
	// Get environment variable
	database, databaseExists := os.LookupEnv(key)
	if !databaseExists {
		logger.VariableNotFound(key)
	}

	return database
}
