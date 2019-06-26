package source

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Komikindo ...
type Komikindo struct{}

// Name ...
func (Komikindo) Name() string { return "komikindo" }

// GetID ...
func (Komikindo) GetID() primitive.ObjectID {
	id, _ := primitive.ObjectIDFromHex("5d0f6dfbe4e1f617cbbe18b6")
	return id
}
