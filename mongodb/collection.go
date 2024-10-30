package mongodb

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

type Order int

const (
	// Ascending order
	Ascending Order = 1

	// Descending order
	Descending Order = -1
)

// OrderInt converts the Order type to an integer
func (o Order) OrderInt() int {
	return int(o)
}

type FieldIndex struct {
	Name  string
	Order Order
}

// NewFieldIndex creates a new field index
func NewFieldIndex(name string, order Order) *FieldIndex {
	return &FieldIndex{Name: name, Order: order}
}

type SingleFieldIndex struct {
	Model *mongo.IndexModel
}

// NewSingleFieldIndex creates a new single field index
func NewSingleFieldIndex(fieldIndex FieldIndex, unique bool) *SingleFieldIndex {
	// Create the index model
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{fieldIndex.Name, fieldIndex.Order.OrderInt()}},
		Options: options.Index().SetUnique(unique),
	}
	return &SingleFieldIndex{Model: &indexModel}
}

// CompoundFieldIndex represents a compound field index
type CompoundFieldIndex struct {
	Model *mongo.IndexModel
}

// NewCompoundFieldIndex creates a new compound field index
func NewCompoundFieldIndex(fieldIndexes []*FieldIndex, unique bool) *CompoundFieldIndex {
	// Create the keys
	keys := bson.D{}
	for _, fieldIndex := range fieldIndexes {
		keys = append(keys, bson.E{Key: fieldIndex.Name, Value: fieldIndex.Order.OrderInt()})
	}

	// Create the index model
	indexModel := mongo.IndexModel{
		Keys:    keys,
		Options: options.Index().SetUnique(unique),
	}
	return &CompoundFieldIndex{Model: &indexModel}
}

// Collection represents a MongoDB collection
type Collection struct {
	Name               string
	SingleFieldIndexes *[]*SingleFieldIndex
	CompoundIndexes    *[]*CompoundFieldIndex
}

// NewCollection creates a new MongoDB collection
func NewCollection(name string, singleFieldIndexes *[]*SingleFieldIndex, compoundIndexes *[]*CompoundFieldIndex) *Collection {
	return &Collection{Name: name, SingleFieldIndexes: singleFieldIndexes, CompoundIndexes: compoundIndexes}
}

// CreateCollection creates the collection
func (c *Collection) CreateCollection(database *mongo.Database) (collection *mongo.Collection, err error) {
	// Get the collection
	collection = database.Collection(c.Name)

	// Create the indexes
	if err = c.createIndexes(collection); err != nil {
		return nil, err
	}

	return collection, nil
}

// CreateIndexes creates the indexes for the collection
func (c *Collection) createIndexes(collection *mongo.Collection) (err error) {
	// Create the single field indexes
	if c.SingleFieldIndexes != nil {
		for _, singleFieldIndex := range *c.SingleFieldIndexes {
			_, err = collection.Indexes().CreateOne(context.Background(), *singleFieldIndex.Model)
			if err != nil {
				return FailedToCreateSingleFieldIndexError
			}
		}
	}

	// Create the compound indexes
	if c.CompoundIndexes != nil {
		for _, compoundIndex := range *c.CompoundIndexes {
			if _, err = collection.Indexes().CreateOne(context.Background(), *compoundIndex.Model); err != nil {
				return FailedToCreateCompoundIndexError
			}
		}
	}
	return nil
}
