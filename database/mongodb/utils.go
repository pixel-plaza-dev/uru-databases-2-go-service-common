package mongodb

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetObjectIdFromString gets the object ID from the string
func GetObjectIdFromString(id string) (*primitive.ObjectID, error) {
	// Check if the ID is empty
	if id == "" {
		return nil, mongo.ErrNoDocuments
	}

	// Create the Object ID from the ID
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return &objectId, nil
}

// PrepareFindOneOptions prepares the find one options
func PrepareFindOneOptions(projection interface{}, sort interface{}) *options.FindOneOptions {
	// Create the find options
	findOptions := options.FindOne()

	// Set the projection
	if projection != nil {
		findOptions.SetProjection(projection)
	}

	// Set the sort
	if sort != nil {
		findOptions.SetSort(sort)
	}

	return findOptions
}
