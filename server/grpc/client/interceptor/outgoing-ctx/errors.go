package outgoing_ctx

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	// FailedToGetOutgoingContextError is the error for failed to get outgoing context
	FailedToGetOutgoingContextError = status.Error(codes.Internal, "failed to get outgoing context")
)
