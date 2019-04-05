package repository

import (
	"time"

	"github.com/bickyeric/arumba/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// IPage ...
type IPage interface {
	FindByEpisode(episodeID, sourceID primitive.ObjectID) (model.Page, error)
	Insert(*model.Page) error
	Update(*model.Page) error
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
	_, err := repo.coll.InsertOne(ctx, page)
	return err
}

func (repo pageRepository) Update(page *model.Page) error {
	page.UpdatedAt = time.Now()
	_, err := repo.coll.UpdateOne(ctx, bson.M{"_id": page.ID}, page)
	return err
}
