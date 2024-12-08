package auth

import (
	"errors"
)

var (
	MissingTokenError        = errors.New("missing token")
	UnexpectedTokenTypeError = errors.New("unexpected token type")
)
