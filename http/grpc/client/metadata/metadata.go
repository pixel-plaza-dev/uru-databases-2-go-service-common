package metadata

import (
	"context"
	commongcloud "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/cloud/gcloud"
	commongrpc "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/http/grpc"
	context2 "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/http/grpc/client/context"
	"google.golang.org/grpc/metadata"
	"strings"
)

type (
	// MetadataField is a field in the metadata
	MetadataField struct {
		Key   string
		Value string
	}

	// CtxMetadata is the metadata for the context
	CtxMetadata struct {
		MetadataFields []MetadataField
	}
)

// NewCtxMetadata creates a new CtxMetadata
func NewCtxMetadata(metadataFields *map[string]string) (*CtxMetadata, error) {
	// Check if the metadata fields are nil
	if metadataFields == nil {
		return nil, context2.NilMetadataFieldsError
	}

	// Add the metadata fields
	var fields []MetadataField
	for key, value := range *metadataFields {
		fields = append(
			fields,
			MetadataField{Key: strings.ToLower(key), Value: value},
		)
	}

	return &CtxMetadata{
		MetadataFields: fields,
	}, nil
}

// NewUnauthenticatedCtxMetadata creates a new unauthenticated CtxMetadata
func NewUnauthenticatedCtxMetadata(gcloudToken string) (*CtxMetadata, error) {
	return NewCtxMetadata(
		&map[string]string{
			commongcloud.AuthorizationMetadataKey: commongrpc.BearerPrefix + " " + gcloudToken,
		},
	)
}

// NewAuthenticatedCtxMetadata creates a new authenticated CtxMetadata
func NewAuthenticatedCtxMetadata(
	gcloudToken string, jwtToken string,
) (*CtxMetadata, error) {
	return NewCtxMetadata(
		&map[string]string{
			commongcloud.AuthorizationMetadataKey: commongrpc.BearerPrefix + " " + gcloudToken,
			commongrpc.AuthorizationMetadataKey:   commongrpc.BearerPrefix + " " + jwtToken,
		},
	)
}

// GetCtxWithMetadata gets the context with the metadata
func GetCtxWithMetadata(
	ctxMetadata *CtxMetadata, ctx context.Context,
) context.Context {
	// Check if the context metadata is nil
	if ctxMetadata == nil {
		return ctx
	}

	// Create metadata
	md := metadata.Pairs()

	// Add the metadata to the context
	for _, field := range ctxMetadata.MetadataFields {
		md.Append(field.Key, field.Value)
	}
	return metadata.NewOutgoingContext(ctx, md)
}

// AppendGCloudTokenToOutgoingContext appends the GCloud token to the outgoing context
func AppendGCloudTokenToOutgoingContext(
	ctx context.Context, gcloudToken string,
) context.Context {
	return metadata.AppendToOutgoingContext(
		ctx,
		commongcloud.AuthorizationMetadataKey,
		commongrpc.BearerPrefix+" "+gcloudToken,
	)
}
