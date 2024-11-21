package redis

import (
	"github.com/go-redis/redis/v8"
	commondatabase "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/database"
	"golang.org/x/net/context"
)

type (
	// ConnectionHandler interface
	ConnectionHandler interface {
		Connect() (*redis.Client, error)
		GetClient() (*redis.Client, error)
		Disconnect()
	}

	// Config struct
	Config struct {
		Uri      string
		Password string
		Database int
	}

	// DefaultConnectionHandler struct
	DefaultConnectionHandler struct {
		Client        *redis.Client
		ClientOptions *redis.Options
	}
)

// NewDefaultConnectionHandler creates a new connection
func NewDefaultConnectionHandler(config *Config) DefaultConnectionHandler {
	// Define the Redis options
	clientOptions := &redis.Options{
		Addr:     config.Uri,
		Password: config.Password,
		DB:       config.Database,
	}

	return DefaultConnectionHandler{
		ClientOptions: clientOptions,
		Client:        nil,
	}
}

// Connect returns a new Redis client
func (d DefaultConnectionHandler) Connect() (*redis.Client, error) {
	// Check if the connection is already established
	if d.Client != nil {
		return d.Client, commondatabase.AlreadyConnectedError
	}

	// Create a new Redis client
	client := redis.NewClient(d.ClientOptions)

	// Ping the Redis server to check the connection
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, commondatabase.FailedToPingError
	}

	// Set client
	d.Client = client

	return client, nil
}

// GetClient returns the Redis client
func (d DefaultConnectionHandler) GetClient() (*redis.Client, error) {
	// Check if the connection is established
	if d.Client == nil {
		return nil, commondatabase.NotConnectedError
	}

	return d.Client, nil
}

// Disconnect closes the Redis client connection
func (d DefaultConnectionHandler) Disconnect() {
	defer func() {
		// Check if the connection is established
		if d.Client == nil {
			return
		}

		// Close the connection
		d.Cancel()
		if err := d.Client.Close(); err != nil {
			panic(commondatabase.FailedToDisconnectError)
		}
	}()
}
