package tls

import (
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc/credentials"
	"os"
)

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

// LoadSystemCredentials loads the system credentials
func LoadSystemCredentials() (credentials.TransportCredentials, error) {
	systemRoots, err := x509.SystemCertPool()
	if err != nil {
		return nil, FailedToLoadSystemCredentialsError
	}
	return credentials.NewTLS(&tls.Config{
		RootCAs: systemRoots,
	}), nil
}
