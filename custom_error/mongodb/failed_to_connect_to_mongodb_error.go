package mongodb

import (
	"github.com/pixel-plaza-dev/uru-databases-2-go-service-common/utils"
	"strings"
)

type FailedToConnectToMongoDbError struct {
	Err error
}

// Error returns a formatted error message for FailedToConnectToMongoDbError
func (l FailedToConnectToMongoDbError) Error() (message string) {
	formattedError := utils.AddBrackets(l.Err.Error())
	return strings.Join([]string{"Failed to connect to MongoDB", formattedError}, " ")
}
