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

// IPage ...
type IPage interface {
	FindByEpisode(episodeID, sourceID primitive.ObjectID) (model.Page, error)
	Insert(*model.Page) error
	GetSources(episodeID primitive.ObjectID) ([]primitive.ObjectID, error)
	Interface
}

type pageRepository struct {
	coll *mongo.Collection
}

// NewPage ...
func NewPage(db *mongo.Database) IPage {
	return pageRepository{db.Collection("pages")}
}

func (repo pageRepository) FindByEpisode(episodeID, sourceID primitive.ObjectID) (model.Page, error) {
	result := model.Page{}
	err := repo.coll.FindOne(ctx, bson.M{"episode_id": episodeID, "source_id": sourceID}).Decode(&result)
	return result, err
}

func (repo pageRepository) Insert(page *model.Page) error {
	page.ID = primitive.NewObjectID()
	page.CreatedAt = time.Now()
	page.UpdatedAt = time.Now()
	_, err := repo.coll.InsertOne(ctx, page)
	return err
}

func (repo pageRepository) GetSources(episodeID primitive.ObjectID) ([]primitive.ObjectID, error) {
	ids := []primitive.ObjectID{}
	cur, err := repo.coll.Find(ctx, bson.M{"episode_id": episodeID})
	if err != nil {
		return ids, err
	}

	c := model.Page{}
	for cur.Next(ctx) {
		if err := cur.Decode(&c); err != nil {
			return ids, err
		}
		ids = append(ids, c.SourceID)
	}

	return ids, err
}

func (repo pageRepository) CreateIndex(ctx context.Context) error {
	models := []mongo.IndexModel{
		{
			Keys:    bson.D{{"episode_id", 1}, {"source_id", 1}},
			Options: options.Index().SetBackground(true).SetUnique(true)},
	}
	return createIndex(ctx, repo.coll, models)
}
