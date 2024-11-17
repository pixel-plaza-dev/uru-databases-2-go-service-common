package context

import (
	"context"
	commongcloud "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/cloud/gcloud"
	commongrpc "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/server/grpc"
	"google.golang.org/grpc/metadata"
	"strings"
)

// MetadataField is a field in the metadata
type MetadataField struct {
	Key   string
	Value string
}

// CtxMetadata is the metadata for the context
type CtxMetadata struct {
	MetadataFields []MetadataField
}

// NewCtxMetadata creates a new CtxMetadata
func NewCtxMetadata(metadataFields map[string]string) *CtxMetadata {
	var fields []MetadataField

	// Add the metadata fields
	for key, value := range metadataFields {
		fields = append(fields, MetadataField{Key: strings.ToLower(key), Value: value})
	}

	return &CtxMetadata{
		MetadataFields: fields,
	}
}

// NewUnauthenticatedCtxMetadata creates a new unauthenticated CtxMetadata
func NewUnauthenticatedCtxMetadata(gcloudToken string) *CtxMetadata {
	return NewCtxMetadata(
		map[string]string{
			commongcloud.AuthorizationMetadataKey: commongrpc.BearerPrefix + " " + gcloudToken,
		},
	)
}

// NewAuthenticatedCtxMetadata creates a new authenticated CtxMetadata
func NewAuthenticatedCtxMetadata(
	gcloudToken string, jwtToken string,
) *CtxMetadata {
	return NewCtxMetadata(
		map[string]string{
			commongcloud.AuthorizationMetadataKey: commongrpc.BearerPrefix + " " + gcloudToken,
			commongrpc.AuthorizationMetadataKey:   commongrpc.BearerPrefix + " " + jwtToken,
		},
	)
}

// GetCtxWithMetadata gets the context with the metadata
func GetCtxWithMetadata(
	ctxMetadata *CtxMetadata, ctx context.Context,
) context.Context {
	// Add the metadata to the context
	for _, field := range ctxMetadata.MetadataFields {
		ctx = metadata.AppendToOutgoingContext(ctx, field.Key, field.Value)
	}
	return ctx
}
