package mongodb

import (
	"github.com/pixel-plaza-dev/uru-databases-2-go-service-common/utils"
	"strings"
)

type FailedToCreateDocumentError struct {
	Err error
}

// Error returns a formatted error message for FailedToCreateDocumentError
func (f FailedToCreateDocumentError) Error() (message string) {
	formattedError := utils.AddBrackets(f.Err.Error())
	return strings.Join([]string{"Failed to create document", formattedError}, " ")
}
