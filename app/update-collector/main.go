package main

import (
	"github.com/bickyeric/arumba"
	"github.com/bickyeric/arumba/connection"
	"github.com/bickyeric/arumba/telegram"
	"github.com/subosito/gotenv"
)

func main() {

	gotenv.Load(".env")

	connection.Connect()
	telegram.ConfigureBot()
	arumba.Configure()

	app := arumba.Instance
	mangacanUpdater := app.InjectMangacanUpdater()

	mangacanUpdater.Run()
	// gocron.Every(1).Minute().Do(updater.Mangacan)

	// <-gocron.Start()
}
