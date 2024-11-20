package listener

import (
	enverror "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/config/env/error"
	"os"
	"strings"
)

// ServicePort is the port of the service
type ServicePort struct {
	Port          string
	FormattedPort string
}

// LoadServicePort load Service port from environment variables
func LoadServicePort(host string, key string) (
	servicePort *ServicePort, err error,
) {
	// Get environment variable
	port, exists := os.LookupEnv(key)
	if !exists {
		return nil, enverror.VariableNotFoundError{Key: key}
	}

	// Build port string
	var portBuilder strings.Builder
	portBuilder.WriteString(host)
	portBuilder.WriteString(":")
	portBuilder.WriteString(port)

	return &ServicePort{
		Port:          port,
		FormattedPort: portBuilder.String(),
	}, nil
}
