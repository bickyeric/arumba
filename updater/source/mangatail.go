package source

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Mangatail ...
type Mangatail struct{}

// Name ...
func (Mangatail) Name() string { return "mangatail" }

// GetID ...
func (Mangatail) GetID() primitive.ObjectID {
	id, _ := primitive.ObjectIDFromHex("5c89e1cb5cff252ae5db8f1e")
	return id
}
