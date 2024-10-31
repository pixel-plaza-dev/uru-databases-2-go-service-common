package listener

import "errors"

var (
	FailedToCloseError  = errors.New("failed to close listener")
	FailedToListenError = errors.New("failed to listen")
	FailedToServeError  = errors.New("failed to serve")
)
