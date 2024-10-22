package mongodb

import (
	customMongoDbError "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/custom_error/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"time"
)

type Config struct {
	Uri     string
	Timeout time.Duration
}

type Connection struct {
	Config *Config
	Client *mongo.Client
	Ctx    context.Context
	Cancel context.CancelFunc
}

// Connect returns a new MongoDB client
func Connect(config *Config) (connection *Connection, err error) {
	// Set client options
	ctx, cancel := context.WithTimeout(context.Background(), config.Timeout)
	clientOptions := options.Client().ApplyURI(config.Uri)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)

	// Create MongoDB Connection struct
	if err != nil {
		return nil, customMongoDbError.FailedToConnectToMongoDbError{Err: err}
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, customMongoDbError.PingToMongoDbFailedError{Err: err}
	}

	return &Connection{Config: config, Client: client, Ctx: ctx, Cancel: cancel}, nil
}

// Disconnect closes the MongoDB client connection
func Disconnect(connection *Connection) {
	defer func() {
		// Close the connection
		connection.Cancel()
		if err := connection.Client.Disconnect(connection.Ctx); err != nil {
			panic(customMongoDbError.FailedToDisconnectFromMongoDbError{Err: err})
		}
	}()
}
