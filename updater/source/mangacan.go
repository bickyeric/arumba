package source

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Mangacan ...
type Mangacan struct{}

// Name ...
func (Mangacan) Name() string { return "mangacan" }

// GetID ...
func (Mangacan) GetID() primitive.ObjectID {
	id, _ := primitive.ObjectIDFromHex("5daddd4b73b1d018e959c85b")
	return id
}
