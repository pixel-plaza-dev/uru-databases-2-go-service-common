package listener

import "errors"

var FailedToCloseError = errors.New("failed to close listener")
var FailedToListenError = errors.New("failed to listen")
var FailedToServeError = errors.New("failed to serve")
