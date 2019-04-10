package source

import "go.mongodb.org/mongo-driver/bson/primitive"

// ISource ...
type ISource interface {
	Name() string
	GetID() primitive.ObjectID
}
