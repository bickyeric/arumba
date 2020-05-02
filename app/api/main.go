package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
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

	basicAuth := apiMiddleware.BasicAuth{Username: os.Getenv("USERNAME"), Password: os.Getenv("PASSWORD")}

	e := echo.New()
	// region    ************************** GLOBAL HTTP MIDDLEWARE **************************
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(apiMiddleware.ErrorHandler)
	// endregion    ************************** GLOBAL HTTP MIDDLEWARE **************************

	r := resolver.New(db)
	config := generated.Config{Resolvers: r}
	config.Directives.Authenticated = basicAuth.IsAuthenticated
	schema := generated.NewExecutableSchema(config)
	graphql := handler.NewDefaultServer(schema)

	kendang := controller.NewKendang(saver)

	// region    ************************** HTTP ROUTER **************************
	e.GET("/", echo.WrapHandler(playground.Handler("GraphQL playground", "/query")))
	e.POST("/query", echo.WrapHandler(graphql), basicAuth.Checker)
	e.POST("/kendang/webhook", kendang.OnHandle)
	// endregion    ************************** HTTP ROUTER **************************

	go e.Start(":1907")

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-signals

	e.Close()
	client.Disconnect(ctx)
}
