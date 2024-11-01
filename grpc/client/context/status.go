package context

import (
	"errors"
	commongrpc "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/grpc"
	"google.golang.org/grpc/status"
)

// ExtractErrorFromStatus extracts the error from the status
func ExtractErrorFromStatus(err error) error {
	st, ok := status.FromError(err)
	if ok {
		return errors.New(st.Message())
	}
	return commongrpc.InternalServerError
}
