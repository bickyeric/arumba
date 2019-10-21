package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/bickyeric/arumba"
	"github.com/bickyeric/arumba/connection"
	"github.com/bickyeric/arumba/controller"
	"github.com/bickyeric/arumba/service/episode"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/subosito/gotenv"
)

func main() {
	ctx := context.Background()
	gotenv.Load(".env")

	client := connection.NewMongo(ctx)
	db := client.Database(os.Getenv("DB_MONGO_DATABASE"))
	app := arumba.New(db)
	saver := episode.NewSaveUpdate(app)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())

	kendang := controller.NewKendang(saver)
	e.POST("/kendang/webhook", kendang.OnHandle)

	e.Logger.Fatal(e.Start(":1907"))

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-signals

	e.Close()
	client.Disconnect(ctx)
}
