package connection

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewMongo() *mongo.Database {
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
		log.Fatal("MongoDB: " + err.Error())
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if err := client.Connect(ctx); err != nil {
		log.Fatal("Failed connecting to MongoDB")
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal("MongoDB is not listening...")
	}

	return client.Database(os.Getenv("DB_MONGO_DATABASE"))
}
