package mongodb

import "errors"

var (
	FailedToConnectError                = errors.New("failed to connect to MongoDB")
	AlreadyConnectedError               = errors.New("connection to MongoDB already established")
	NotConnectedError                   = errors.New("connection to MongoDB not established")
	FailedToDisconnectError             = errors.New("failed to disconnect from MongoDB")
	FailedToCreateDocumentError         = errors.New("failed to create document")
	FailedToStartSessionError           = errors.New("failed to start session")
	FailedToPingError                   = errors.New("failed to ping MongoDB")
	FailedToCreateSingleFieldIndexError = errors.New("failed to create single field index")
	FailedToCreateCompoundIndexError    = errors.New("failed to create compound index")
)
