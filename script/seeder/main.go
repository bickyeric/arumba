package main

import (
	"context"
	"fmt"
	"os"

	"github.com/bickyeric/arumba/connection"
	"github.com/bickyeric/arumba/model"
	"github.com/bickyeric/arumba/repository"
	log "github.com/sirupsen/logrus"
	"github.com/subosito/gotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var sources = []struct {
	ID       string
	Name     string
	Hostname string
}{
	{"5d0f6dede4e1f617cbbe1865", "komikcast", "https://komikcast.co.id/"},
	{"5d0f6dfbe4e1f617cbbe18b6", "komikindo", "https://www.komikindo.web.id/"},
	{"5daddd4b73b1d018e959c85b", "mangacan", "http://www.mangacanblog.com/"},
	{"5d13989eaddc0b6d19eef333", "mangaku", "https://mangaku.in/"},
	{"5c89e1cb5cff252ae5db8f1e", "mangatail", "https://www.mangatail.me/"},
	{"5e1f31a200832f65a5e44826", "komiku", "https://komiku.co.id/"},
	{"5e6f9ad81ae41533caac7d56", "mangakyo", "https://www.mangakyo.com/"},
}

func main() {
	gotenv.Load(".env")
	log.SetFormatter(&log.JSONFormatter{})

	db := connection.NewMongo(context.Background()).Database(os.Getenv("DB_MONGO_DATABASE"))
	sourceSeeder(db)
	comicSeeder(db)
}

func sourceSeeder(db *mongo.Database) {
	repo := repository.NewSource(db)
	for _, s := range sources {
		id, _ := primitive.ObjectIDFromHex(s.ID)
		source := model.Source{
			ID:       id,
			Name:     s.Name,
			Hostname: s.Hostname,
		}

		err := repo.Insert(&source)
		if err != nil {
			log.Warn(err)
		}
	}
}

func comicSeeder(db *mongo.Database) {
	repo := repository.NewComic(db)
	for i := 1; i <= 100; i++ {
		comic := model.Comic{
			Name: fmt.Sprintf("Comic %d", i),
		}
		err := repo.Insert(&comic)
		if err != nil {
			log.Warn(err)
		}
	}
}
