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
	Insert(s model.Source) error
	All() ([]*model.Source, error)
}

type sourceRepository struct {
	coll *mongo.Collection
}

// NewSource ...
func NewSource(db *mongo.Database) ISource {
	return sourceRepository{db.Collection("sources")}
}

func (repo sourceRepository) All() ([]*model.Source, error) {
	sources := []*model.Source{}
	cur, err := repo.coll.Find(ctx, bson.M{})
	if err != nil {
		return sources, err
	}

	for cur.Next(ctx) {
		s := model.Source{}
		if err := cur.Decode(&s); err != nil {
			return sources, err
		}
		sources = append(sources, &s)
	}

	return sources, nil
}

func (repo sourceRepository) Insert(s model.Source) error {
	_, err := repo.coll.InsertOne(ctx, s)
	return err
}

func (repo sourceRepository) FindByID(id primitive.ObjectID) (source model.Source, err error) {
	err = repo.coll.FindOne(ctx, bson.M{"_id": id}).Decode(&source)
	return source, err
}
