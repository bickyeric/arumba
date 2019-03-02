package arumba

import (
	"database/sql"

	"github.com/bickyeric/arumba/repository"
)

// New ...
func New(db *sql.DB) Arumba {
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
