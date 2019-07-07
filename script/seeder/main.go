package main

import (
	"time"

	"github.com/bickyeric/arumba/connection"
	"github.com/bickyeric/arumba/model"
	"github.com/bickyeric/arumba/repository"
	"github.com/bickyeric/arumba/updater"
	log "github.com/sirupsen/logrus"
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load(".env")
	log.SetFormatter(&log.JSONFormatter{})

	db := connection.NewMongo()

	repo := repository.NewSource(db)
	for _, s := range updater.Sources {
		source := model.Source{
			ID:        s.GetID(),
			Name:      s.Name(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		err := repo.Insert(source)
		if err != nil {
			log.Warn(err)
		}
	}
}
