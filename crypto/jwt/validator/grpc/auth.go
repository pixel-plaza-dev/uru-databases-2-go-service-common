package grpc

import (
	"context"
	commonredisauth "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/database/redis/auth"
	pbauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/compiled/auth"
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
		authClient          *pbauth.AuthClient
		redisTokenValidator commonredisauth.TokenValidator
	}
)

// NewDefaultTokenValidator creates a new default token validator
func NewDefaultTokenValidator(
	authTokenSource *oauth.TokenSource,
	authClient *pbauth.AuthClient,
	redisTokenValidator commonredisauth.TokenValidator,
) (*DefaultTokenValidator, error) {
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
		_, err := (*d.authClient).IsRefreshTokenValid(
			context.Background(), &pbauth.IsRefreshTokenValidRequest{
				JwtId: jwtId,
			},
		)
		return err != nil, err
	}

	// Check if the access token is valid
	_, err := (*d.authClient).IsAccessTokenValid(
		context.Background(), &pbauth.IsAccessTokenValidRequest{
			JwtId: jwtId,
		},
	)
	return err != nil, err
}
