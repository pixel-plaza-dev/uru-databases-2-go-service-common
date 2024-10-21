package mongodb

import (
	"github.com/pixel-plaza-dev/uru-databases-2-go-service-common/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"time"
)

// Connect returns a new MongoDB client
func Connect(uri string, logger *logger.MongoDbLogger, timeout time.Duration) (*mongo.Client, context.Context, context.CancelFunc, error) {
	// Set client options
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return client, ctx, cancel, err
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return client, ctx, cancel, err
	}

	// Log the connection
	logger.ConnectedToMongoDB()

	return client, ctx, cancel, nil
}

// Disconnect closes the MongoDB client connection
func Disconnect(client *mongo.Client, context context.Context, cancel context.CancelFunc, logger *logger.MongoDbLogger) {
	defer func() {
		// Close the connection
		cancel()
		if err := client.Disconnect(context); err != nil {
			logger.FailedToDisconnectFromMongoDb(err)
		}

		// Log the disconnection
		logger.DisconnectedFromMongoDB()
	}()
}
