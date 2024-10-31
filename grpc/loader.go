package grpc

import (
	enverror "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/env/error"
	"os"
	"strings"
)

// Uri is the URI for the gRPC service
type Uri struct {
	Host string
	Port string
	Uri  string
}

// LoadUri loads the URI from the environment
func LoadUri(hostKey string, portKey string) (*Uri, error) {
	// Get host from environment
	host, exists := os.LookupEnv(hostKey)
	if !exists {
		return nil, enverror.VariableNotFoundError{Key: hostKey}
	}

	// Get port from environment
	port, exists := os.LookupEnv(portKey)
	if !exists {
		return nil, enverror.VariableNotFoundError{Key: portKey}
	}

	// Build URI string
	var uriBuilder strings.Builder
	uriBuilder.WriteString(host)
	uriBuilder.WriteString(":")
	uriBuilder.WriteString(port)

	return &Uri{
		Host: host,
		Port: port,
		Uri:  uriBuilder.String(),
	}, nil
}
