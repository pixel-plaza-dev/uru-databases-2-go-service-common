package database

import "errors"

var (
	AlreadyConnectedError   = errors.New("connection to database already established")
	FailedToPingError       = errors.New("failed to ping database")
	NotConnectedError       = errors.New("connection to database not established")
	FailedToDisconnectError = errors.New("failed to disconnect from database")
)
