package arumba

import (
	"github.com/bickyeric/arumba/repository"
)

// Arumba ...
type Arumba struct {
	SourceRepo  repository.ISource
	ComicRepo   repository.IComic
	EpisodeRepo repository.IEpisode
	PageRepo    repository.IPage
}
