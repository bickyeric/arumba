package repository

//go:generate mockgen -destination mock/page.go -package=mock -source page.go

import (
	"context"
	"time"

	"github.com/bickyeric/arumba/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// IPage ...
type IPage interface {
	FindByEpisodeSource(episodeID, sourceID primitive.ObjectID) (model.Page, error)
	FindByEpisode(ctx context.Context, episodeID primitive.ObjectID, first, offset int) ([]model.Page, error)
	Insert(*model.Page) error
	CreateIndex(context.Context) error
}

type pageRepository struct {
	coll *mongo.Collection
}

// NewPage ...
func NewPage(db *mongo.Database) IPage {
	return pageRepository{db.Collection("pages")}
}

func (repo pageRepository) FindByEpisode(ctx context.Context, episodeID primitive.ObjectID, first, offset int) ([]model.Page, error) {
	var pages []model.Page
	cur, err := repo.coll.Find(ctx,
		primitive.M{"episode_id": episodeID},
		options.Find().SetLimit(int64(first)).SetSkip(int64(offset)),
	)
	if err != nil {
		return pages, err
	}

	for cur.Next(ctx) {
		p := model.Page{}
		if err := cur.Decode(&p); err != nil {
			return pages, err
		}
		pages = append(pages, p)
	}
	return pages, nil
}

func (repo pageRepository) FindByEpisodeSource(episodeID, sourceID primitive.ObjectID) (model.Page, error) {
	result := model.Page{}
	err := repo.coll.FindOne(ctx, primitive.M{"episode_id": episodeID, "source_id": sourceID}).Decode(&result)
	return result, err
}

func (repo pageRepository) Insert(page *model.Page) error {
	page.ID = primitive.NewObjectID()
	page.CreatedAt = time.Now()
	page.UpdatedAt = time.Now()
	_, err := repo.coll.InsertOne(ctx, page)
	return err
}

func (repo pageRepository) CreateIndex(ctx context.Context) error {
	models := []mongo.IndexModel{
		{
			Keys:    primitive.D{{"episode_id", 1}, {"source_id", 1}},
			Options: options.Index().SetBackground(true).SetUnique(true)},
	}
	return createIndex(ctx, repo.coll, models)
}
