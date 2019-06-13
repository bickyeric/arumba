package repository

import (
	"time"

	"github.com/bickyeric/arumba/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// IComic ...
type IComic interface {
	Find(name string) (model.Comic, error)
	FindByID(id primitive.ObjectID) (model.Comic, error)
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
	comic.ID = primitive.NewObjectID()
	comic.CreatedAt = time.Now()
	_, err := repo.coll.InsertOne(ctx, comic)
	return err
}

func (repo comicRepository) FindByID(id primitive.ObjectID) (c model.Comic, err error) {
	err = repo.coll.FindOne(ctx, bson.M{"_id": id}).Decode(&c)
	return c, err
}

func (repo comicRepository) Find(name string) (c model.Comic, err error) {
	err = repo.coll.FindOne(ctx, bson.M{"name": bson.M{"$regex": ".*" + name + ".*", "$options": "i"}}).Decode(&c)
	return c, err
}

func (repo comicRepository) FindAll(name string) ([]model.Comic, error) {
	comics := []model.Comic{}
	cur, err := repo.coll.Find(ctx,
		bson.M{"name": bson.M{"$regex": ".*" + name + ".*", "$options": "i"}},
		options.Find().SetLimit(5),
	)
	if err != nil {
		return comics, err
	}

	c := model.Comic{}
	for cur.Next(ctx) {
		if err := cur.Decode(&c); err != nil {
			return comics, err
		}
		comics = append(comics, c)
	}

	return comics, nil
}
