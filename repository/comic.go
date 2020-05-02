package repository

//go:generate mockgen -destination mock/comic.go -package=mock -source comic.go

import (
	"context"
	"time"

	"github.com/bickyeric/arumba/external"
	"github.com/bickyeric/arumba/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// IComic ...
type IComic interface {
	Find(name string) (model.Comic, error)
	FindByID(id primitive.ObjectID) (model.Comic, error)
	FindByName(context.Context, string) (model.Comic, error)
	FindAll(ctx context.Context, name string, first, offset int) ([]model.Comic, error)
	Insert(*model.Comic) error
	CreateIndex(context.Context) error
}

type comicRepository struct {
	coll *mongo.Collection
}

// NewComic ...
func NewComic(db external.MongoDatabase) IComic {
	return comicRepository{db.Collection("comics")}
}

func (repo comicRepository) Insert(comic *model.Comic) error {
	comic.ID = primitive.NewObjectID()
	comic.CreatedAt = time.Now()
	comic.UpdatedAt = comic.CreatedAt
	_, err := repo.coll.InsertOne(ctx, comic)
	return err
}

func (repo comicRepository) FindByID(id primitive.ObjectID) (c model.Comic, err error) {
	err = repo.coll.FindOne(ctx, primitive.M{"_id": id}).Decode(&c)
	return c, err
}

func (repo comicRepository) Find(name string) (c model.Comic, err error) {
	err = repo.coll.FindOne(ctx, primitive.M{"name": name}).Decode(&c)
	return c, err
}

func (repo comicRepository) FindByName(ctx context.Context, name string) (c model.Comic, err error) {
	err = repo.coll.FindOne(ctx, primitive.M{"name": name}).Decode(&c)
	return c, err
}

func (repo comicRepository) FindAll(ctx context.Context, name string, first, offset int) ([]model.Comic, error) {
	var comics []model.Comic
	cur, err := repo.coll.Find(ctx,
		primitive.M{"name": primitive.M{"$regex": ".*" + name + ".*", "$options": "i"}},
		options.Find().SetLimit(int64(first)).SetSkip(int64(offset)),
	)
	if err != nil {
		return comics, err
	}

	for cur.Next(ctx) {
		c := model.Comic{}
		if err := cur.Decode(&c); err != nil {
			return comics, err
		}
		comics = append(comics, c)
	}

	return comics, nil
}

func (repo comicRepository) CreateIndex(ctx context.Context) error {
	models := []mongo.IndexModel{
		{
			Keys:    primitive.D{{"name", 1}},
			Options: options.Index().SetName("name").SetBackground(true).SetUnique(true),
		},
	}
	return createIndex(ctx, repo.coll, models)
}
