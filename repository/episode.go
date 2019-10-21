package repository

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
	Count(comicID primitive.ObjectID, bound ...float64) (int, error)
	No(comicID primitive.ObjectID, offset int, bound ...float64) (float64, error)
	FindByNo(comicID primitive.ObjectID, no int) (*model.Episode, error)
	Insert(*model.Episode) error
}

type episodeRepository struct {
	coll *mongo.Collection
}

// NewEpisode ...
func NewEpisode(db *mongo.Database) IEpisode {
	return episodeRepository{db.Collection("episodes")}
}

func (repo episodeRepository) Count(comicID primitive.ObjectID, bound ...float64) (int, error) {
	filter := bson.M{"comic_id": comicID}
	if len(bound) == 2 {
		filter["no"] = bson.M{
			"$gte": bound[0],
			"$lte": bound[1],
		}
	}
	totalEpisode, err := repo.coll.CountDocuments(ctx, filter)
	return int(totalEpisode), err
}

func (repo episodeRepository) No(comicID primitive.ObjectID, offset int, bound ...float64) (float64, error) {
	ep := model.Episode{}
	filter := bson.M{"comic_id": comicID}
	if len(bound) > 0 {
		filter["no"] = bson.M{
			"$gte": bound[0],
			"$lte": bound[1],
		}
	}
	res := repo.coll.FindOne(ctx, filter,
		options.FindOne().SetSort(bson.M{"no": 1}).SetSkip(int64(offset)))
	err := res.Decode(&ep)
	return 0, err
}

func (repo episodeRepository) Insert(ep *model.Episode) error {
	ep.ID = primitive.NewObjectID()
	ep.CreatedAt = time.Now()
	_, err := repo.coll.InsertOne(ctx, ep)
	return err
}

func (repo episodeRepository) FindByNo(comicID primitive.ObjectID, no int) (*model.Episode, error) {
	ep := new(model.Episode)
	err := repo.coll.FindOne(ctx, bson.M{"comic_id": comicID, "no": no}).Decode(ep)
	return ep, err
}
