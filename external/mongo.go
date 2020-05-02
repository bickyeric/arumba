package external

//go:generate mockgen -destination mock/mongo.go -package=mock -source mongo.go

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDatabase is abstraction for *mongo.Database
type MongoDatabase interface {
	Collection(name string, opts ...*options.CollectionOptions) *mongo.Collection
}
