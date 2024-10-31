package context

import (
	commongrpc "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/grpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

// MetadataField is a field in the metadata
type MetadataField struct {
	Key   string
	Value string
}

// CtxMetadata is the metadata for the context
type CtxMetadata struct {
	Token MetadataField
}

// NewCtxMetadata creates a new CtxMetadata
func NewCtxMetadata(token string) *CtxMetadata {
	return &CtxMetadata{
		Token: MetadataField{Key: commongrpc.AuthorizationHeaderKey, Value: token},
	}
}

// GetCtxWithMetadata gets the context with the metadata
func GetCtxWithMetadata(ctxMetadata *CtxMetadata, ctx context.Context) context.Context {
	// Add the metadata to the context
	for _, field := range []MetadataField{ctxMetadata.Token} {
		ctx = metadata.AppendToOutgoingContext(ctx, field.Key, field.Value)
	}
	return ctx
}
