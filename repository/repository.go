package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Interface interface {
	CreateIndex(context.Context) error
}

func createIndex(ctx context.Context, coll *mongo.Collection, models []mongo.IndexModel) error {
	iv := coll.Indexes()
	_, err := iv.CreateMany(context.Background(), models, options.CreateIndexes().SetMaxTime(2*time.Second))
	return err
}
