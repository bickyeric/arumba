package repository

import (
	"context"

	"github.com/bickyeric/arumba/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// IComic ...
type IComic interface {
	Find(name string) (model.Comic, error)
	FindAll(name string) ([]model.Comic, error)
	Insert(*model.Comic) error
}

type comicRepository struct {
	coll *mongo.Collection
}

// NewComic ...
func NewComic(db *mongo.Database) IComic {
	return comicRepository{db.Collection("comics")}
}

func (repo comicRepository) Insert(comic *model.Comic) error {
	_, err := repo.coll.InsertOne(context.Background(), comic)
	return err
}

func (repo comicRepository) Find(name string) (model.Comic, error) {
	c := model.Comic{}
	err := repo.coll.FindOne(context.Background(), bson.M{"name": "/.*" + name + ".*/"}).Decode(&c)
	return c, err
}

func (repo comicRepository) FindAll(name string) ([]model.Comic, error) {
	comics := []model.Comic{}
	cur, err := repo.coll.Find(context.Background(), bson.M{"name": "/.*" + name + ".*/"})

	c := model.Comic{}
	for cur.Next(context.Background()) {
		if err := cur.Decode(&c); err != nil {
			return comics, err
		}
	}

	return comics, err
}
