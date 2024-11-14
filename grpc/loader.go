package grpc

import (
	"crypto/tls"
	"crypto/x509"
	enverror "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/env/error"
	"golang.org/x/net/context"
	"google.golang.org/api/idtoken"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
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

// LoadTLSCredentials loads the TLS credentials
func LoadTLSCredentials(pemServerCAPath string) (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := os.ReadFile(pemServerCAPath)
	if err != nil {
		return nil, err
	}

	// Create a certificate pool from the certificate
	certPool := x509.NewCertPool()

	// Append the certificates from the PEM file
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, FailedToAddCAPemError
	}

	// Create the credentials and return it
	config := &tls.Config{
		RootCAs: certPool,
	}

	return credentials.NewTLS(config), nil
}

// LoadServiceAccountCredentials loads the service account credentials
func LoadServiceAccountCredentials(ctx context.Context, url string) (*oauth.TokenSource, error) {
	// Use the default service account credentials
	ts, err := idtoken.NewTokenSource(ctx, url)
	if err != nil {
		return nil, FailedToCreateTokenSourceError
	}

	return &oauth.TokenSource{TokenSource: ts}, nil
}
