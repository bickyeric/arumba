package connection

import (
	"context"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewMongo ...
func NewMongo(ctx context.Context) *mongo.Client {
	opts := options.Client().
		SetHosts(strings.Split(os.Getenv("DB_MONGO_HOST"), ","))

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err = client.Ping(ctx, nil); err != nil {
		panic("MongoDB is not listening...")
	}

	return client
}
