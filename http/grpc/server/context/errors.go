package context

import "errors"

var (
	MissingTokenInContextError              = errors.New("missing token in context")
	UnexpectedTokenTypeInContextError       = errors.New("unexpected type in context")
	MissingTokenClaimsInContextError        = errors.New("missing token claims in context")
	UnexpectedTokenClaimsTypeInContextError = errors.New("unexpected token claims type in context")
)
