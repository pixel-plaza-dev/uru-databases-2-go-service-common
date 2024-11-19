package auth

import (
	"context"
	commonvalidator "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/crypto/jwt/validator"
	commongrpc "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/server/grpc"
	grpcctx "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/server/grpc/server/context"
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
		validator          commonvalidator.Validator
		methodsToIntercept map[string]bool
	}
)

// NewInterceptor creates a new authentication interceptor
func NewInterceptor(
	validator commonvalidator.Validator, methodsToIntercept map[string]bool,
) *Interceptor {
	return &Interceptor{
		validator:          validator,
		methodsToIntercept: methodsToIntercept,
	}
}

// GetTokenFromMetadata gets the token from the metadata
func (i *Interceptor) GetTokenFromMetadata(md metadata.MD) (string, error) {
	// Get the authorization from the metadata
	authorization := md.Get(commongrpc.AuthorizationMetadataKey)
	tokenIdx := commongrpc.TokenIdx.Int()
	if len(authorization) <= tokenIdx {
		return "", AuthorizationMetadataNotProvidedError
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

// GetMethodName gets the method name from the full method
func (i *Interceptor) GetMethodName(fullMethod string) string {
	parts := strings.Split(fullMethod, "/")
	if len(parts) < 3 {
		return ""
	}
	return parts[2]
}

// Authenticate returns the authentication interceptor
func (i *Interceptor) Authenticate(mustBeRefreshToken bool) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		// Get method name
		methodName := i.GetMethodName(info.FullMethod)

		// Check if the method should be intercepted
		intercept, ok := i.methodsToIntercept[methodName]
		if !intercept || !ok {
			return handler(ctx, req)
		}

		// Get metadata from the context
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, MetadataNotProvidedError
		}

		// Get the token from the metadata
		tokenString, err := i.GetTokenFromMetadata(md)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}

		// Validate the token and get the validated claims
		claims, err := i.validator.GetValidatedClaims(tokenString, mustBeRefreshToken)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}

		// Set the claims in the context
		ctx = grpcctx.SetCtxClaims(&ctx, claims)

		return handler(ctx, req)
	}
}
