package context

import (
	"errors"
)

var (
	NilMetadataFieldsError = errors.New("metadata fields cannot be nil")
	NilCtxMetadataError    = errors.New("context metadata cannot be nil")
)
