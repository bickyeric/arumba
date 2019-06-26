package source

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Mangaku ...
type Mangaku struct{}

// Name ...
func (Mangaku) Name() string { return "mangaku" }

// GetID ...
func (Mangaku) GetID() primitive.ObjectID {
	id, _ := primitive.ObjectIDFromHex("5d13989eaddc0b6d19eef333")
	return id
}
