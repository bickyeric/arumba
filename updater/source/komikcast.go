package source

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Komikcast ...
type Komikcast struct{}

// Name ...
func (Komikcast) Name() string { return "komikcast" }

// GetID ...
func (Komikcast) GetID() primitive.ObjectID {
	id, _ := primitive.ObjectIDFromHex("5d0f6dede4e1f617cbbe1865")
	return id
}
