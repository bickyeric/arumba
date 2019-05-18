package repository

import (
	"github.com/bickyeric/arumba/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ISource ...
type ISource interface {
	FindByID(id primitive.ObjectID) (model.Source, error)
}

type sourceRepository struct {
	coll *mongo.Collection
}

// NewSource ...
func NewSource(db *mongo.Database) ISource {
	return sourceRepository{db.Collection("sources")}
}

func (repo sourceRepository) FindByID(id primitive.ObjectID) (source model.Source, err error) {
	err = repo.coll.FindOne(ctx, bson.M{"_id": id}).Decode(&source)
	return source, err
}
