package auth

import (
	commonvalidator "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/crypto/jwt/validator"
	commongrpc "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/server/grpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
)

// Interceptor is the interceptor for the authentication
type Interceptor struct {
	validator          commonvalidator.Validator
	methodsToIntercept map[string]bool
}

// NewInterceptor creates a new authentication interceptor
func NewInterceptor(validator commonvalidator.Validator, methodsToIntercept map[string]bool) *Interceptor {
	return &Interceptor{
		validator:          validator,
		methodsToIntercept: methodsToIntercept,
	}
}

// GetTokenFromMetadata gets the token from the metadata
func GetTokenFromMetadata(md metadata.MD) (string, error) {
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
	if len(authorizationFields) != 2 {
		return "", commongrpc.AuthorizationMetadataInvalidError
	}

	// Check the first field
	if authorizationFields[0] != commongrpc.BearerPrefix {
		return "", commongrpc.MissingBearerPrefixError
	}

	return authorizationFields[1], nil
}

// UnaryServerInterceptor return the interceptor function
func (i *Interceptor) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// Check if the method should be intercepted
		intercept, ok := i.methodsToIntercept[info.FullMethod]
		if !intercept || !ok {
			return handler(ctx, req)
		}

		// Get metadata from the context
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, MetadataNotProvidedError
		}

		// Get the token from the metadata
		tokenString, err := GetTokenFromMetadata(md)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}

		// Validate the token and get the claims
		claims, err := i.validator.GetClaims(tokenString)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}

		// Set the claims in the context
		ctx = commongrpc.SetContextClaims(&ctx, claims)

		return handler(ctx, req)
	}
}
