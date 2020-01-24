package main

import (
	"context"
	"os"

	"github.com/bickyeric/arumba/connection"
	"github.com/bickyeric/arumba/model"
	"github.com/bickyeric/arumba/repository"
	log "github.com/sirupsen/logrus"
	"github.com/subosito/gotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var sources = []struct {
	ID   string
	Name string
}{
	{"5d0f6dede4e1f617cbbe1865", "komikcast"},
	{"5d0f6dfbe4e1f617cbbe18b6", "komikindo"},
	{"5daddd4b73b1d018e959c85b", "mangacan"},
	{"5d13989eaddc0b6d19eef333", "mangaku"},
	{"5c89e1cb5cff252ae5db8f1e", "mangatail"},
}

func main() {
	gotenv.Load(".env")
	log.SetFormatter(&log.JSONFormatter{})

	db := connection.NewMongo(context.Background()).Database(os.Getenv("DB_MONGO_DATABASE"))

	repo := repository.NewSource(db)
	for _, s := range sources {
		id, _ := primitive.ObjectIDFromHex(s.ID)
		source := model.Source{
			ID:   id,
			Name: s.Name,
		}

		err := repo.Insert(source)
		if err != nil {
			log.Warn(err)
		}
	}
}
