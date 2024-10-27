package error

type FailedToCreateDocumentError struct {
	Err error
}

// Error returns a formatted error message for FailedToCreateDocumentError
func (f FailedToCreateDocumentError) Error() (message string) {
	return "Failed to create document: " + f.Err.Error()
}
