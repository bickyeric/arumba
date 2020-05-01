package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/99designs/gqlgen/handler"
	apiMiddleware "github.com/bickyeric/arumba/api/middleware"
	"github.com/bickyeric/arumba/connection"
	"github.com/bickyeric/arumba/controller"
	"github.com/bickyeric/arumba/generated"
	"github.com/bickyeric/arumba/repository"
	"github.com/bickyeric/arumba/resolver"
	"github.com/bickyeric/arumba/service/episode"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load(".env")
	ctx := context.Background()

	// region    ************************** CONNECTION **************************
	client := connection.NewMongo(ctx)
	db := client.Database(os.Getenv("DB_MONGO_DATABASE"))
	// endregion    ************************** CONNECTION **************************

	// region    ************************** REPO **************************
	sourceRepo := repository.NewSource(db)
	comicRepo := repository.NewComic(db)
	episodeRepo := repository.NewEpisode(db)
	pageRepo := repository.NewPage(db)
	// endregion    ************************** REPO **************************

	// region    ************************** SERVICE **************************
	saver := episode.NewSaveUpdate(sourceRepo, comicRepo, episodeRepo, pageRepo)
	// endregion    ************************** SERVICE **************************

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	basicAuth := apiMiddleware.BasicAuth{Username: os.Getenv("USERNAME"), Password: os.Getenv("PASSWORD")}
	authConfig := middleware.DefaultBasicAuthConfig
	authConfig.Validator = basicAuth.Assignor

	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(middleware.BasicAuthWithConfig(authConfig))
	e.Use(apiMiddleware.ErrorHandler)

	kendang := controller.NewKendang(saver)

	episode := resolver.NewEpisode(pageRepo)
	r := resolver.New(episode, db)
	config := generated.Config{Resolvers: r}
	config.Directives.Authenticated = basicAuth.IsAuthenticated

	schema := generated.NewExecutableSchema(config)
	e.GET("/", echo.WrapHandler(handler.Playground("GraphQL playground", "/query")))
	e.POST("/query", echo.WrapHandler(handler.GraphQL(schema)))
	e.POST("/kendang/webhook", kendang.OnHandle)

	go e.Start(":1907")

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-signals

	e.Close()
	client.Disconnect(ctx)
}
