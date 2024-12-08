package metadata

import (
	commongcloud "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/cloud/gcloud"
	commongrpc "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/http/grpc"
	"google.golang.org/grpc/metadata"
	"strings"
)

// GetTokenFromMetadata gets the token from the metadata
func GetTokenFromMetadata(md metadata.MD, tokenKey string) (string, error) {
	// Get the authorization from the metadata
	authorization := md.Get(tokenKey)
	tokenIdx := commongrpc.TokenIdx.Int()
	if len(authorization) <= tokenIdx {
		return "", commongrpc.AuthorizationMetadataNotProvidedError
	}

	// Get the authorization value from the metadata
	authorizationValue := authorization[tokenIdx]

	// Split the authorization value by space
	authorizationFields := strings.Split(authorizationValue, " ")

	// Check if the authorization value is valid
	if len(authorizationFields) != 2 || authorizationFields[0] != commongrpc.BearerPrefix {
		return "", commongrpc.AuthorizationMetadataInvalidError
	}

	return authorizationFields[1], nil
}

// GetAuthorizationTokenFromMetadata gets the authorization token from the metadata
func GetAuthorizationTokenFromMetadata(md metadata.MD) (string, error) {
	return GetTokenFromMetadata(md, commongrpc.AuthorizationMetadataKey)
}

// GetGCloudAuthorizationTokenFromMetadata gets the GCloud authorization token from the metadata
func GetGCloudAuthorizationTokenFromMetadata(md metadata.MD) (string, error) {
	return GetTokenFromMetadata(md, commongcloud.AuthorizationMetadataKey)
}
