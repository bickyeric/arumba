package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/bickyeric/arumba"
	"github.com/bickyeric/arumba/connection"
	"github.com/bickyeric/arumba/generated"
	"github.com/bickyeric/arumba/repository"
	"github.com/subosito/gotenv"
)

const defaultPort = "8080"

func main() {
	gotenv.Load(".env")

	db := connection.NewMongo()

	ComicRepo := repository.NewComic(db)
	EpisodeRepo := repository.NewEpisode(db)
	SourceRepo := repository.NewSource(db)

	resolver := &arumba.Resolver{
		ComicRepo:   ComicRepo,
		SourceRepo:  SourceRepo,
		EpisodeRepo: EpisodeRepo,
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: resolver})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
