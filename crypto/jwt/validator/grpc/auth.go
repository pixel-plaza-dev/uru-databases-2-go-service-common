package grpc

import (
	"context"
	commongcloud "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/cloud/gcloud"
	commonredisauth "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/database/redis/auth"
	commongrpcclient "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/http/grpc/client"
	pbauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/auth"
	"google.golang.org/grpc/credentials/oauth"
)

type (
	// TokenValidator interface
	TokenValidator interface {
		IsTokenValid(token string, jwtId string, isRefreshToken bool) (bool, error)
	}

	// DefaultTokenValidator struct
	DefaultTokenValidator struct {
		accessToken         string
		authClient          pbauth.AuthClient
		redisTokenValidator commonredisauth.TokenValidator
	}
)

// NewDefaultTokenValidator creates a new default token validator
func NewDefaultTokenValidator(
	authTokenSource *oauth.TokenSource,
	authClient pbauth.AuthClient,
	redisTokenValidator commonredisauth.TokenValidator,
) (*DefaultTokenValidator, error) {
	// Check if either the token source or the auth client is nil
	if authTokenSource == nil {
		return nil, commongcloud.NilTokenSourceError
	}
	if authClient == nil {
		return nil, commongrpcclient.NilClientError
	}

	// Get the token from the token source
	token, err := authTokenSource.Token()
	if err != nil {
		return nil, err
	}

	return &DefaultTokenValidator{
		authClient:          authClient,
		accessToken:         token.AccessToken,
		redisTokenValidator: redisTokenValidator,
	}, nil
}

// IsTokenValid checks if the token is valid
func (d *DefaultTokenValidator) IsTokenValid(
	token string,
	jwtId string,
	isRefreshToken bool,
) (bool, error) {
	// Check if redis is enabled
	if d.redisTokenValidator != nil {
		// Check if the token is valid
		return d.redisTokenValidator.IsTokenValid(token)
	}

	// Validate the token
	if isRefreshToken {
		// Check if the refresh token is valid
		response, err := d.authClient.IsRefreshTokenValid(
			context.Background(), &pbauth.IsRefreshTokenValidRequest{
				JwtId: jwtId,
			},
		)
		if err != nil {
			return false, err
		}
		return response.IsValid, nil
	}

	// Check if the access token is valid
	response, err := d.authClient.IsAccessTokenValid(
		context.Background(), &pbauth.IsAccessTokenValidRequest{
			JwtId: jwtId,
		},
	)
	if err != nil {
		return false, err
	}
	return response.IsValid, nil
}
