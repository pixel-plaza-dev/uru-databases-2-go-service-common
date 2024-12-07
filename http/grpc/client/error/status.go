package error

import (
	"errors"
	commonflag "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/config/flag"
	commongrpc "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/http/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ExtractErrorFromStatus extracts the error from the status
func ExtractErrorFromStatus(mode *commonflag.ModeFlag, err error) error {
	// Check if the flag mode is nil
	if mode == nil {
		return commonflag.NilModeFlagError
	}

	st, ok := status.FromError(err)

	// Check if the error is a status error
	if !ok {
		// Check the flag mode
		if mode.IsProd() {
			return commongrpc.InternalServerError
		}
		return err
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
