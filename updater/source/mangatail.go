package source

import (
	"github.com/bickyeric/arumba/updater"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Mangatail ...
type Mangatail struct{}

var _ updater.ISource = (*Mangatail)(nil)

// Name ...
func (Mangatail) Name() string { return "mangatail" }

// GetID ...
func (Mangatail) GetID() primitive.ObjectID {
	id, _ := primitive.ObjectIDFromHex("5c9511f561a8d04fa844b666")
	return id
}
