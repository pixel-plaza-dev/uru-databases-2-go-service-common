package shop

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BusinessClient is the MongoDB business client model
type BusinessClient struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	BusinessID primitive.ObjectID `json:"business_id" bson:"business_id"`
	ClientID   primitive.ObjectID `json:"client_id" bson:"client_id"`
}

// BusinessOwner is the MongoDB business owner model
type BusinessOwner struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	BusinessID primitive.ObjectID `json:"business_id" bson:"business_id"`
	OwnerID    primitive.ObjectID `json:"owner_id" bson:"owner_id"`
	JoinedAt   primitive.DateTime `json:"joined_at" bson:"joined_at"`
	RemovedAt  primitive.DateTime `json:"removed_at,omitempty" bson:"removed_at,omitempty"`
	Percentage float64            `json:"percentage,omitempty" bson:"percentage,omitempty"`
}

// Business is the MongoDB business model
type Business struct {
	ID                 primitive.ObjectID   `json:"id" bson:"_id"`
	MarketCategoriesID []primitive.ObjectID `json:"market_categories_id" bson:"market_categories_id"`
	ProfilePicture     primitive.ObjectID   `json:"profile_picture,omitempty" bson:"profile_picture,omitempty"`
	Name               string               `json:"name" bson:"name"`
	Description        string               `json:"description" bson:"description"`
	JoinedAt           primitive.DateTime   `json:"joined_at" bson:"joined_at"`
	RemovedAt          primitive.DateTime   `json:"removed_at,omitempty" bson:"removed_at,omitempty"`
	AdminRevisions     []primitive.ObjectID `json:"admin_revision,omitempty" bson:"admin_revision,omitempty"`
	IsSuspended        bool                 `json:"is_suspended,omitempty" bson:"is_suspended,omitempty"`
}

// MarketCategory is the MongoDB market category model
type MarketCategory struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
}

// AdminRevision is the MongoDB admin revision model
type AdminRevision struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	AdminID     primitive.ObjectID `json:"admin_id" bson:"admin_id"`
	IsSuspended bool               `json:"is_suspended,omitempty" bson:"is_suspended,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	StartedAt   primitive.DateTime `json:"started_at" bson:"started_at"`
	EndedAt     primitive.DateTime `json:"ended_at,omitempty" bson:"ended_at,omitempty"`
}

// Book is the MongoDB book model
type Book struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	ISBN        string             `json:"isbn" bson:"isbn"`
	Author      string             `json:"author" bson:"author"`
	Publisher   string             `json:"publisher" bson:"publisher"`
	Genres      []string           `json:"genres" bson:"genres"`
	PublishedAt primitive.DateTime `json:"published_at" bson:"published_at"`
	Pages       int64              `json:"pages" bson:"pages"`
	Language    string             `json:"language" bson:"language"`
}

// Clothing is the MongoDB clothing model
type Clothing struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Size     string             `json:"size" bson:"size"`
	Color    string             `json:"color" bson:"color"`
	Material string             `json:"material" bson:"material"`
	Gender   string             `json:"gender,omitempty" bson:"gender,omitempty"`
	Season   string             `json:"season,omitempty" bson:"season,omitempty"`
}
