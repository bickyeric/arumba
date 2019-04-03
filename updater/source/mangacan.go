package source

import (
	"github.com/bickyeric/arumba/updater"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Mangacan ...
type Mangacan struct{}

var _ updater.ISource = (*Mangacan)(nil)

// Name ...
func (Mangacan) Name() string { return "mangacan" }

// GetID ...
func (Mangacan) GetID() primitive.ObjectID {
	id, _ := primitive.ObjectIDFromHex("5c9511f561a8d04fa844b666")
	return id
}
