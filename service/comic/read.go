package comic

import (
	"math/rand"
	"time"

	"github.com/bickyeric/arumba/model"
	"github.com/bickyeric/arumba/repository"
)

type Read struct {
	ComicRepo   repository.IComic
	EpisodeRepo repository.IEpisode
	PageRepo    repository.IPage
}

func (r Read) Perform(comicName string, episodeNo float64) ([]*model.Page, error) {
	comic, err := r.ComicRepo.FindByName(comicName)
	if err != nil {
		return nil, err
	}

	episode, err := r.EpisodeRepo.FindByNo(comic.ID, episodeNo)
	if err != nil {
		return nil, err
	}

	sources := r.EpisodeRepo.GetSources(episode.ID)

	rand.Seed(time.Now().Unix())
	n := rand.Int() % len(sources)

	return r.PageRepo.FindByEpisode(episode.ID, sources[n])
}
