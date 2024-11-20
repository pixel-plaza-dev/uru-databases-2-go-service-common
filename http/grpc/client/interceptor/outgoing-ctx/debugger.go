package outgoing_ctx

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type (
	// OutgoingCtxDebugger interface
	OutgoingCtxDebugger interface {
		PrintOutgoingCtx() grpc.UnaryClientInterceptor
	}

	// Interceptor is the interceptor for the debug
	Interceptor struct {
		logger Logger
	}
)

// NewInterceptor creates a new debug interceptor
func NewInterceptor(logger Logger) *Interceptor {
	return &Interceptor{
		logger: logger,
	}
}

// PrintOutgoingCtx prints the outgoing context
func (i *Interceptor) PrintOutgoingCtx() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		// Get the outgoing context
		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			return FailedToGetOutgoingContextError
		}

		// Print the metadata
		for key, values := range md {
			for _, value := range values {
				i.logger.LogKeyValue(key, value)
			}
		}

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
