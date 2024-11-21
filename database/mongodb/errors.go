package mongodb

import "errors"

var (
	FailedToCreateDocumentError         = errors.New("failed to create document")
	FailedToStartSessionError           = errors.New("failed to start session")
	FailedToCreateSingleFieldIndexError = errors.New("failed to create single field index")
	FailedToCreateCompoundIndexError    = errors.New("failed to create compound index")
)
