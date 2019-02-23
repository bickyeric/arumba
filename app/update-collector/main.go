package main

import (
	"github.com/bickyeric/arumba"
	"github.com/bickyeric/arumba/connection"
	"github.com/bickyeric/arumba/telegram"
	"github.com/bickyeric/arumba/updater/source"
	"github.com/subosito/gotenv"
)

func main() {

	gotenv.Load(".env")

	connection.Connect()
	telegram.ConfigureBot()
	arumba.Configure()

	app := arumba.Instance
	updater := app.InjectUpdateRunner()

	mangacan := source.Mangacan{}
	updater.Run(mangacan)
	// gocron.Every(1).Minute().Do(updater.Run, mangacan)

	mangatail := source.Mangatail{}
	updater.Run(mangatail)

	// <-gocron.Start()
}
