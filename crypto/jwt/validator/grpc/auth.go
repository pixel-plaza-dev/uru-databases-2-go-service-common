package grpc

import (
	"context"
	commongrpcclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/http/grpc/client/context"
	pbauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/compiled/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/oauth"
)

type (
	// TokenValidator interface
	TokenValidator interface {
		IsTokenValid(token string, jwtId string, isRefreshToken bool) bool
	}

	// DefaultTokenValidator struct
	DefaultTokenValidator struct {
		accessToken string
		authClient  *pbauth.AuthClient
	}
)

// NewDefaultTokenValidator creates a new default token validator
func NewDefaultTokenValidator(
	authTokenSource *oauth.TokenSource,
	authClient *pbauth.AuthClient,
) (*DefaultTokenValidator, error) {
	// Get the token from the token source
	token, err := authTokenSource.Token()
	if err != nil {
		return nil, err
	}

	return &DefaultTokenValidator{
		authClient:  authClient,
		accessToken: token.AccessToken,
	}, nil
}

// IsTokenValid checks if the token is valid
func (d *DefaultTokenValidator) IsTokenValid(
	token string,
	jwtId string,
	isRefreshToken bool,
) bool {
	// Get context metadata
	var ctxMetadata *commongrpcclientctx.CtxMetadata
	if jwtId != "" {
		ctxMetadata = commongrpcclientctx.NewAuthenticatedCtxMetadata(
			d.accessToken,
			token,
		)
	} else {
		ctxMetadata = commongrpcclientctx.NewUnauthenticatedCtxMetadata(d.accessToken)
	}

	// Get outgoing context
	grpcCtx := commongrpcclientctx.GetCtxWithMetadata(
		ctxMetadata,
		context.Background(),
	)

	// Validate the token
	if isRefreshToken {
		// Check if the refresh token is valid
		response, err := (*d.authClient).IsRefreshTokenValid(
			grpcCtx, &pbauth.IsRefreshTokenValidRequest{
				JwtId: jwtId,
			},
		)
		if err != nil {
			return false
		}
		return response.Code == uint32(codes.OK)

	}

	// Check if the access token is valid
	response, err := (*d.authClient).IsAccessTokenValid(
		grpcCtx, &pbauth.IsAccessTokenValidRequest{
			JwtId: jwtId,
		},
	)
	if err != nil {
		return false
	}
	return response.Code == uint32(codes.OK)
}