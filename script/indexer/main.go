package main

import (
	"context"
	"os"

	"github.com/bickyeric/arumba/connection"
	"github.com/bickyeric/arumba/repository"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load(".env")
}

func main() {
	ctx := context.Background()
	client := connection.NewMongo(ctx)
	db := client.Database(os.Getenv("DB_MONGO_DATABASE"))

	var repo repository.Interface
	repo = repository.NewComic(db)
	repo.CreateIndex(ctx)

	repo = repository.NewEpisode(db)
	repo.CreateIndex(ctx)

	repo = repository.NewPage(db)
	repo.CreateIndex(ctx)
}
