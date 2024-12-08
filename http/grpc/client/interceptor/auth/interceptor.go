package auth

import (
	"context"
	commongcloud "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/cloud/gcloud"
	commongrpc "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/http/grpc"
	commongrpcclientmd "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/http/grpc/client/metadata"
	commongrpcinfo "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/http/grpc/info"
	commongrpcmd "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/http/grpc/metadata"
	pbtypesgrpc "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/oauth"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Interceptor is the interceptor for the authentication
type Interceptor struct {
	accessToken       string
	grpcInterceptions *map[pbtypesgrpc.Method]pbtypesgrpc.Interception
}

// NewInterceptor creates a new authentication interceptor
func NewInterceptor(
	tokenSource *oauth.TokenSource,
	grpcInterceptions *map[pbtypesgrpc.Method]pbtypesgrpc.Interception,
) (*Interceptor, error) {
	// Check if the token source is nil
	if tokenSource == nil {
		return nil, commongcloud.NilTokenSourceError
	}

	// Get the access token from the token source
	token, err := tokenSource.Token()
	if err != nil {
		return nil, err
	}

	// Check if the gRPC interceptions is nil
	if grpcInterceptions == nil {
		return nil, commongrpc.NilGRPCInterceptionsError
	}

	return &Interceptor{
		accessToken: token.AccessToken,
	}, nil
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
	) (err error) {
		// Get the method name
		methodName := commongrpcinfo.GetMethodName(method)

		// Check if the method should be intercepted
		var ctxMetadata *commongrpcclientmd.CtxMetadata
		interception, ok := (*i.grpcInterceptions)[pbtypesgrpc.NewMethod(
			methodName,
		)]
		if !ok || interception == pbtypesgrpc.None {
			// Create the unauthenticated context metadata
			ctxMetadata, err = commongrpcclientmd.NewUnauthenticatedCtxMetadata(i.accessToken)
		} else {
			// Get metadata from the context
			md, ok := metadata.FromIncomingContext(ctx)
			if !ok {
				return status.Error(codes.Unauthenticated, commongrpc.MissingMetadataError.Error())
			}

			// Get the token from the metadata
			tokenString, err := commongrpcmd.GetAuthorizationTokenFromMetadata(md)
			if err != nil {
				return status.Error(codes.Unauthenticated, err.Error())
			}

			// Create the authenticated context metadata
			ctxMetadata, err = commongrpcclientmd.NewAuthenticatedCtxMetadata(i.accessToken, tokenString)
		}

		// Check if there was an error
		if err != nil {
			return status.Error(codes.Aborted, err.Error())
		}

		// Get the gRPC client context with the metadata
		ctx = commongrpcclientmd.GetCtxWithMetadata(ctxMetadata, ctx)

		// Invoke the original invoker
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
