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
	Count(comicID int) (int, error)
	No(comicID, offset int) (float64, error)
	FindByNo(comicID primitive.ObjectID, no float64) (*model.Episode, error)
	GetSources(episodeID primitive.ObjectID) []primitive.ObjectID
	Insert(*model.Episode) error
}

type episodeRepository struct {
	coll *mongo.Collection
}

// NewEpisode ...
func NewEpisode(db *mongo.Database) IEpisode {
	return episodeRepository{db.Collection("episodes")}
}

func (repo episodeRepository) Count(comicID int) (int, error) {
	totalEpisode, err := repo.coll.CountDocuments(ctx, bson.M{"comic_id": comicID})
	return int(totalEpisode), err
}

func (repo episodeRepository) No(comicID, offset int) (float64, error) {
	ep := model.Episode{}
	res := repo.coll.FindOne(ctx, bson.M{"comic_id": comicID},
		options.FindOne().SetSkip(int64(offset)))
	err := res.Decode(&ep)
	return ep.No, err
}

func (repo episodeRepository) Insert(ep *model.Episode) error {
	ep.ID = primitive.NewObjectID()
	ep.CreatedAt = time.Now()
	_, err := repo.coll.InsertOne(ctx, ep)
	return err
}

func (repo episodeRepository) FindByNo(comicID primitive.ObjectID, no float64) (*model.Episode, error) {
	ep := new(model.Episode)
	err := repo.coll.FindOne(ctx, bson.M{"comic_id": comicID, "no": no}).Decode(ep)
	return ep, err
}

func (repo episodeRepository) GetSources(episodeID primitive.ObjectID) []primitive.ObjectID {
	sourceIds := []primitive.ObjectID{}
	// rows, err := repo.Query("SELECT source_id FROM episode_source WHERE episode_id=?", episodeID)
	// if err != nil {
	// 	return sourceIds
	// }
	// for rows.Next() {
	// 	var id int
	// 	rows.Scan(&id)
	// 	sourceIds = append(sourceIds, id)
	// }
	return sourceIds
}
