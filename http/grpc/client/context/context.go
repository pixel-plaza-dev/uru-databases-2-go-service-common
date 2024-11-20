package context

import (
	"context"
	"errors"
	commongrpc "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/http/grpc"
	commongrpcserverctx "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/http/grpc/server/context"
	"google.golang.org/grpc/metadata"
)

// GetOutgoingCtx returns a context with the token string
func GetOutgoingCtx(ctx context.Context) (context.Context, error) {
	// Get the token string from the context
	token, err := commongrpcserverctx.GetCtxTokenString(ctx)
	if err != nil {
		// Check if the token is missing
		if errors.Is(err, commongrpcserverctx.MissingTokenInContextError) {
			return context.Background(), nil
		}
		return nil, err
	}

	// Append the token to the gRPC context
	grpcCtx := metadata.AppendToOutgoingContext(context.Background(), commongrpc.AuthorizationMetadataKey, token)

	return grpcCtx, nil
}
