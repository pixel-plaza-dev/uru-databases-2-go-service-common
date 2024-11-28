package context

import "errors"

var (
	MissingTokenError                   = errors.New("missing token")
	UnexpectedTokenTypeError            = errors.New("unexpected type")
	MissingTokenClaimsError             = errors.New("missing token claims")
	MissingTokenClaimsUserIdError       = errors.New("missing token claims user id")
	MissingTokenClaimsSharedUserIdError = errors.New("missing token claims shared user id")
	MissingTokenClaimsIdError           = errors.New("missing token claims id")
	UnexpectedTokenClaimsTypeError      = errors.New("unexpected token claims type")
	FailedToGetPeerFromContextError     = errors.New("failed to get peer from context")
)
