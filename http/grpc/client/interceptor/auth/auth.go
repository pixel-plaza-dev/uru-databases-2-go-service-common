package auth

import (
	"context"
	"errors"
	commongrpc "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/http/grpc"
	commongrpcctx "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/http/grpc/client/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/oauth"
)

type (
	// Authentication interface
	Authentication interface {
		Authenticate() grpc.UnaryClientInterceptor
	}

	// Interceptor is the interceptor for the authentication
	Interceptor struct {
		accessToken string
	}
)

// NewInterceptor creates a new authentication interceptor
func NewInterceptor(tokenSource *oauth.TokenSource) (*Interceptor, error) {
	// Get the access token from the token source
	token, err := tokenSource.Token()
	if err != nil {
		return nil, err
	}

	return &Interceptor{
		accessToken: token.AccessToken,
	}, nil
}

// GetCtxTokenString tries to get the token string from the context metadata of the gRPC request
func (i *Interceptor) GetCtxTokenString(ctx context.Context) (string, error) {
	// Get the token from the context
	value := ctx.Value(commongrpc.AuthorizationMetadataKey)
	if value == nil {
		return "", MissingTokenError
	}

	// Check the type of the value
	token, ok := value.(string)
	if !ok {
		return "", UnexpectedTokenTypeError
	}

	return token, nil
}

// Authenticate returns a new unary client interceptor that adds authentication metadata to the context
func (i *Interceptor) Authenticate() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		var ctxMetadata *commongrpcctx.CtxMetadata

		// Get the JWT token
		jwtToken, err := i.GetCtxTokenString(ctx)
		if err == nil {
			// Create the authenticated context metadata
			ctxMetadata = commongrpcctx.NewAuthenticatedCtxMetadata(
				i.accessToken, jwtToken,
			)
		} else {
			// Check if the error is a missing token error
			if errors.Is(err, MissingTokenError) {
				// Create the unauthenticated context metadata
				ctxMetadata = commongrpcctx.NewUnauthenticatedCtxMetadata(i.accessToken)
			} else {
				return err
			}
		}

		// Get the gRPC client context with the metadata
		ctx = commongrpcctx.GetCtxWithMetadata(ctxMetadata, ctx)

		// Invoke the original invoker
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
