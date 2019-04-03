package arumba

import (
	"github.com/bickyeric/arumba/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

// New ...
func New(db *mongo.Database) Arumba {
	return Arumba{
		ComicRepo:   repository.NewComic(db),
		EpisodeRepo: repository.NewEpisode(db),
		PageRepo:    repository.NewPage(db),
	}
}

type Arumba struct {
	ComicRepo   repository.IComic
	EpisodeRepo repository.IEpisode
	PageRepo    repository.IPage
}
