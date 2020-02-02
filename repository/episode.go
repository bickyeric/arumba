package repository

//go:generate mockgen -destination mock/episode.go -package=mock -source episode.go

import (
	"context"
	"time"

	"github.com/bickyeric/arumba/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()

// IEpisode ...
type IEpisode interface {
	FindByNo(comicID primitive.ObjectID, no int) (*model.Episode, error)
	FindAll(ctx context.Context, comicID primitive.ObjectID, first, offset int) ([]model.Episode, error)
	Insert(*model.Episode) error
	CreateIndex(context.Context) error
}

type episodeRepository struct {
	coll *mongo.Collection
}

// NewEpisode ...
func NewEpisode(db *mongo.Database) IEpisode {
	return episodeRepository{db.Collection("episodes")}
}

func (repo episodeRepository) FindAll(ctx context.Context, comicID primitive.ObjectID, first, offset int) ([]model.Episode, error) {
	var episodes []model.Episode
	cur, err := repo.coll.Find(ctx,
		bson.M{"comic_id": comicID},
		options.Find().SetLimit(int64(first)).SetSkip(int64(offset)),
	)
	if err != nil {
		return episodes, err
	}
	err = cur.All(ctx, &episodes)
	return episodes, err
}

func (repo episodeRepository) Insert(ep *model.Episode) error {
	ep.ID = primitive.NewObjectID()
	ep.CreatedAt = time.Now()
	ep.UpdatedAt = ep.CreatedAt
	_, err := repo.coll.InsertOne(ctx, ep)
	return err
}

func (repo episodeRepository) FindByNo(comicID primitive.ObjectID, no int) (*model.Episode, error) {
	ep := new(model.Episode)
	err := repo.coll.FindOne(ctx, bson.M{"comic_id": comicID, "no": no}).Decode(ep)
	return ep, err
}

func (repo episodeRepository) CreateIndex(ctx context.Context) error {
	models := []mongo.IndexModel{
		{
			Keys:    bson.D{{"comic_id", 1}, {"no", 1}},
			Options: options.Index().SetBackground(true).SetUnique(true)},
	}
	return createIndex(ctx, repo.coll, models)
}
