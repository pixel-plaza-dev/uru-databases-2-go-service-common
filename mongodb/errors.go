package mongodb

import "errors"

var FailedToConnectError = errors.New("failed to connect to MongoDB")
var AlreadyConnectedError = errors.New("connection to MongoDB already established")
var NotConnectedError = errors.New("connection to MongoDB not established")
var FailedToDisconnectError = errors.New("failed to disconnect from MongoDB")
var FailedToCreateDocumentError = errors.New("failed to create document")
var FailedToStartSessionError = errors.New("failed to start session")
var FailedToPingError = errors.New("failed to ping MongoDB")
var FailedToCreateSingleFieldIndexError = errors.New("failed to create single field index")
var FailedToCreateCompoundIndexError = errors.New("failed to create compound index")
