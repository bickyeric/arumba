package main

import (
	"github.com/bickyeric/arumba"
	"github.com/bickyeric/arumba/connection"
	"github.com/bickyeric/arumba/updater"
	"github.com/bickyeric/arumba/updater/source"
	"github.com/jasonlvhit/gocron"
	log "github.com/sirupsen/logrus"
	"github.com/subosito/gotenv"
)

func main() {

	gotenv.Load(".env")
	log.SetFormatter(&log.JSONFormatter{})

	db := connection.NewMongo()

	bot := arumba.NewBot()
	kendang := connection.NewKendang()

	app := arumba.New(db)
	updater := updater.NewRunner(bot, kendang, app)

	// updater.Run(source.Mangacan{})

	gocron.Every(1).Minute().Do(updater.Run, source.Mangacan{})
	gocron.Every(1).Minute().Do(updater.Run, source.Mangatail{})

	// updater.Run(source.Mangatail{})

	<-gocron.Start()
}
