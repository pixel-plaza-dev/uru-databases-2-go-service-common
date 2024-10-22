package mongodb

import (
	"github.com/pixel-plaza-dev/uru-databases-2-go-service-common/utils"
	"strings"
)

type FailedToDisconnectFromMongoDbError struct {
	Err error
}

// Error returns a formatted error message for FailedToDisconnectFromMongoDbError
func (l FailedToDisconnectFromMongoDbError) Error() (message string) {
	formattedError := utils.AddBrackets(l.Err.Error())
	return strings.Join([]string{"Failed to disconnect from MongoDB", formattedError}, " ")
}
