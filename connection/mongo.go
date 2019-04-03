package connection

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewMongo() (*mongo.Database, error) {
	uri := ""
	if os.Getenv("DB_MONGO_USERNAME") != "" {
		uri = fmt.Sprintf("mongodb://%s:%s@%s",
			os.Getenv("DB_MONGO_USERNAME"),
			os.Getenv("DB_MONGO_PASSWORD"),
			os.Getenv("DB_MONGO_HOST"))
	} else {
		uri = fmt.Sprintf("mongodb://%s",
			os.Getenv("DB_MONGO_HOST"))
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if err := client.Connect(ctx); err != nil {
		return nil, err
	}

	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	return client.Database(os.Getenv("DB_MONGO_DATABASE")), nil
}
