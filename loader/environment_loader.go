package loader

import (
	customEnvironmentError "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/custom_error/environment"
	"os"
	"strings"
)

type ServicePort struct {
	Port          string
	FormattedPort string
}

// LoadServicePort load Service port from environment variables
func LoadServicePort(key string) (servicePort *ServicePort, err error) {
	// Get environment variable
	port, exists := os.LookupEnv(key)
	if !exists {
		return nil, customEnvironmentError.VariableNotFoundError{VariableName: key}
	}

	// Build port string
	var portBuilder strings.Builder
	portBuilder.WriteString(":")
	portBuilder.WriteString(port)

	return &ServicePort{
		Port:          port,
		FormattedPort: portBuilder.String(),
	}, nil
}

// LoadMongoDBURI load MongoDB URI from environment variables
func LoadMongoDBURI(key string) (uri string, err error) {
	// Get environment variable
	uri, uriExists := os.LookupEnv(key)
	if !uriExists {
		return "", customEnvironmentError.VariableNotFoundError{VariableName: key}
	}
	return uri, nil
}

// LoadMongoDBName load MongoDB database name from environment variables
func LoadMongoDBName(key string) (database string, err error) {
	// Get environment variable
	database, databaseExists := os.LookupEnv(key)
	if !databaseExists {
		return "", customEnvironmentError.VariableNotFoundError{VariableName: key}
	}

	return database, nil
}
