package main

import (
	"context"
	"os"

	"github.com/bickyeric/arumba"
	"github.com/bickyeric/arumba/connection"
	"github.com/bickyeric/arumba/updater"
	"github.com/bickyeric/arumba/updater/source"
	"github.com/jasonlvhit/gocron"
	log "github.com/sirupsen/logrus"
	"github.com/subosito/gotenv"
)

func main() {
	ctx := context.Background()

	gotenv.Load(".env")
	log.SetFormatter(&log.JSONFormatter{})

	db := connection.NewMongo(ctx).Database(os.Getenv("DB_MONGO_DATABASE"))

	bot := arumba.NewBot()
	kendang := connection.NewKendang()

	app := arumba.New(db)
	updateRunner := updater.NewRunner(bot, kendang, app)

	for _, s := range updater.Sources {
		gocron.Every(1).Minute().Do(updateRunner.Run, s)
	}
	updateRunner.Run(source.Mangaku{})

	<-gocron.Start()
}
