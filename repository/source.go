package repository

import (
	"time"

	"github.com/bickyeric/arumba/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ISource ...
type ISource interface {
	FindByID(id primitive.ObjectID) (model.Source, error)
	Insert(s *model.Source) error
}

type sourceRepository struct {
	coll *mongo.Collection
}

// NewSource ...
func NewSource(db *mongo.Database) ISource {
	return sourceRepository{db.Collection("sources")}
}

func (repo sourceRepository) Insert(s *model.Source) error {
	s.ID = primitive.NewObjectID()
	s.CreatedAt = time.Now()
	s.UpdatedAt = s.CreatedAt
	_, err := repo.coll.InsertOne(ctx, s)
	return err
}

func (repo sourceRepository) FindByID(id primitive.ObjectID) (source model.Source, err error) {
	err = repo.coll.FindOne(ctx, primitive.M{"_id": id}).Decode(&source)
	return source, err
}
