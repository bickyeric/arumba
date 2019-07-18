package connection

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewMongo ...
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

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("MongoDB: " + err.Error())
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		log.Fatal("MongoDB is not listening...")
	}

	return client.Database(os.Getenv("DB_MONGO_DATABASE"))
}
