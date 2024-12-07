package outgoing_ctx

import (
	"google.golang.org/grpc"
)

// OutgoingCtxDebugger interface
type OutgoingCtxDebugger interface {
	PrintOutgoingCtx() grpc.UnaryClientInterceptor
}
