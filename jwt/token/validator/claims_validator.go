package validator

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	pbauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/auth"
	gogrpcclient "github.com/ralvarezdev/go-grpc/client"
	gojwt "github.com/ralvarezdev/go-jwt"
	gojwtredisauth "github.com/ralvarezdev/go-jwt/redis/auth"
	gojwtinterception "github.com/ralvarezdev/go-jwt/token/interception"
	goloadergcloud "github.com/ralvarezdev/go-loader/cloud/gcloud"
	"google.golang.org/grpc/credentials/oauth"
)

// DefaultClaimsValidator struct
type DefaultClaimsValidator struct {
	accessToken         string
	authClient          pbauth.AuthClient
	redisTokenValidator gojwtredisauth.TokenValidator
}

// NewDefaultClaimsValidator creates a new default claims validator
func NewDefaultClaimsValidator(
	authTokenSource *oauth.TokenSource,
	authClient pbauth.AuthClient,
	redisTokenValidator gojwtredisauth.TokenValidator,
) (*DefaultClaimsValidator, error) {
	// Check if either the token source or the auth client is nil
	if authTokenSource == nil {
		return nil, goloadergcloud.NilTokenSourceError
	}
	if authClient == nil {
		return nil, gogrpcclient.NilClientError
	}

	// Get the token from the token source
	token, err := authTokenSource.Token()
	if err != nil {
		return nil, err
	}

	return &DefaultClaimsValidator{
		authClient:          authClient,
		accessToken:         token.AccessToken,
		redisTokenValidator: redisTokenValidator,
	}, nil
}

// ValidateClaims validates the claims
func (d *DefaultClaimsValidator) ValidateClaims(
	claims *jwt.MapClaims,
	interception gojwtinterception.Interception,
) (bool, error) {
	// Check if is a refresh token
	isRefreshToken, ok := (*claims)[gojwt.IsRefreshTokenClaim].(bool)
	if !ok {
		return false, IsRefreshTokenClaimNotValidError
	}

	// Get the JWT Identifier
	jwtId, ok := (*claims)[gojwt.IdClaim].(string)
	if !ok {
		return false, IdClaimNotValidError
	}

	// Check if it must be refresh token
	if !isRefreshToken && interception == gojwtinterception.RefreshToken {
		return false, MustBeRefreshTokenError
	}

	// Check if it must be access token
	if isRefreshToken && interception == gojwtinterception.AccessToken {
		return false, MustBeAccessTokenError
	}

	// Check if redis is enabled
	if d.redisTokenValidator != nil {
		// Check if the token is valid
		return d.redisTokenValidator.IsTokenValid(jwtId)
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
