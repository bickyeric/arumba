package main

import (
	"github.com/bickyeric/arumba"
	"github.com/bickyeric/arumba/connection"
	"github.com/bickyeric/arumba/service/telegraph"
	"github.com/bickyeric/arumba/updater"
	"github.com/bickyeric/arumba/updater/source"
	log "github.com/sirupsen/logrus"
	"github.com/subosito/gotenv"
)

func main() {

	gotenv.Load(".env")
	log.SetFormatter(&log.JSONFormatter{})

	db, err := connection.NewMongo()
	if err != nil {
		log.Fatal(err)
	}

	bot := arumba.NewBot()
	kendang := connection.NewKendang()
	telegraphPageCreator := telegraph.NewCreatePage()

	app := arumba.New(db)
	updater := updater.NewRunner(bot, kendang, app, telegraphPageCreator)

	updater.Run(source.Mangacan{})

	// gocron.Every(1).Minute().Do(updater.Run, mangacan)

	// updater.Run(source.Mangatail{})

	// <-gocron.Start()
}
