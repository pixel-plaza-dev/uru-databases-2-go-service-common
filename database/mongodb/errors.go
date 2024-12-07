package mongodb

import "errors"

var (
	FailedToCreateDocumentError         = errors.New("failed to create document")
	FailedToStartSessionError           = errors.New("failed to start session")
	FailedToCreateSingleFieldIndexError = errors.New("failed to create single field index")
	FailedToCreateCompoundIndexError    = errors.New("failed to create compound index")
	NilConfigError                      = errors.New("mongodb connection config cannot be nil")
	NilClientError                      = errors.New("mongodb client cannot be nil")
)
