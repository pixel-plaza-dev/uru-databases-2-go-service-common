package mongodb

import (
	"github.com/pixel-plaza-dev/uru-databases-2-go-service-common/utils"
	"strings"
)

type PingToMongoDbFailedError struct {
	Err error
}

// Error returns a formatted error message for PingToMongoDbFailedError
func (l PingToMongoDbFailedError) Error() (message string) {
	formattedError := utils.AddBrackets(l.Err.Error())
	return strings.Join([]string{"Ping to MongoDB failed", formattedError}, " ")
}
