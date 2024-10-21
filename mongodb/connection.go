package mongodb

import (
	"github.com/pixel-plaza-dev/uru-databases-2-go-service-common/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

// Connect returns a new MongoDB client
func Connect(uri string, logger logger.MongoDbLogger) (*mongo.Client, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		logger.ErrorConnectingToMongoDB(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		logger.ErrorConnectingToMongoDB(err)
	}

	return client, nil
}
