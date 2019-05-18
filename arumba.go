package arumba

import (
	"github.com/bickyeric/arumba/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

// New ...
func New(db *mongo.Database) Arumba {
	return Arumba{
		SourceRepo:  repository.NewSource(db),
		ComicRepo:   repository.NewComic(db),
		EpisodeRepo: repository.NewEpisode(db),
		PageRepo:    repository.NewPage(db),
	}
}

// Arumba ...
type Arumba struct {
	SourceRepo  repository.ISource
	ComicRepo   repository.IComic
	EpisodeRepo repository.IEpisode
	PageRepo    repository.IPage
}
