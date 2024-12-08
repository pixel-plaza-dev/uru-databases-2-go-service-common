package context

import (
	"errors"
)

var (
	MetadataNotFoundError    = errors.New("metadata not found in context")
	MissingTokenError        = errors.New("missing token")
	UnexpectedTokenTypeError = errors.New("unexpected token type")
	NilMetadataFieldsError   = errors.New("metadata fields cannot be nil")
)
