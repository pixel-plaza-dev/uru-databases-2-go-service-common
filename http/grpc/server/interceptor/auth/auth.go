package auth

import (
	"context"
	commonvalidator "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/crypto/jwt/validator"
	commongrpcserverctx "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/http/grpc/server/context"
	pbtypes "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
)

type (
	// Authentication interface
	Authentication interface {
		Authenticate() grpc.UnaryServerInterceptor
	}

	// Interceptor is the interceptor for the authentication
	Interceptor struct {
		validator         commonvalidator.Validator
		grpcInterceptions *map[pbtypes.GRPCMethod]pbtypes.Interception
	}
)

// NewInterceptor creates a new authentication interceptor
func NewInterceptor(
	validator commonvalidator.Validator,
	grpcInterceptions *map[pbtypes.GRPCMethod]pbtypes.Interception,
) (*Interceptor, error) {
	// Check if gRPC interceptions is nil
	if grpcInterceptions == nil {
		return nil, GRPCInterceptionsNilError
	}

	return &Interceptor{
		validator:         validator,
		grpcInterceptions: grpcInterceptions,
	}, nil
}

// GetMethodName gets the method name from the full method
func (i *Interceptor) GetMethodName(fullMethod string) string {
	parts := strings.Split(fullMethod, "/")
	if len(parts) < 3 {
		return ""
	}
	return parts[2]
}

// Authenticate returns the authentication interceptor
func (i *Interceptor) Authenticate() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		// Get method name
		methodName := i.GetMethodName(info.FullMethod)

		// Check if the method should be intercepted
		interception, ok := (*i.grpcInterceptions)[pbtypes.NewGRPCMethod(
			methodName,
		)]
		if !ok || interception == pbtypes.None {
			return handler(ctx, req)
		}

		// Get metadata from the context
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, MetadataNotProvidedError
		}

		// Get the token from the metadata
		tokenString, err := commongrpcserverctx.GetAuthorizationTokenFromMetadata(md)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}

		// Validate the token and get the validated claims
		claims, err := i.validator.GetValidatedClaims(tokenString, interception)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}

		// Set the token string and token claims to the context
		ctx = commongrpcserverctx.SetCtxTokenString(ctx, tokenString)
		ctx = commongrpcserverctx.SetCtxTokenClaims(ctx, claims)

		return handler(ctx, req)
	}
}