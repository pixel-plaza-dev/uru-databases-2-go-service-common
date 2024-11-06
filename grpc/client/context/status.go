package context

import (
	"errors"
	commongrpc "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ExtractErrorFromStatus extracts the error from the status
func ExtractErrorFromStatus(err error) error {
	st, ok := status.FromError(err)

	// Check if the error is a status error
	if !ok {
		return commongrpc.InternalServerError
	}

	// Check the code
	code := st.Code()

	switch code {
	case codes.Unavailable:
		return commongrpc.ServiceUnavailable
	case codes.Unauthenticated:
		return commongrpc.Unauthenticated
	case codes.PermissionDenied:
		return commongrpc.Unauthorized
	default:
		return errors.New(st.Message())
	}
}
