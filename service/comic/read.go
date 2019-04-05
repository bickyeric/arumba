package comic

import (
	"math/rand"
	"time"

	"github.com/bickyeric/arumba/connection"
	"github.com/bickyeric/arumba/model"
	"github.com/bickyeric/arumba/repository"
)

// Read ...
type Read struct {
	ComicRepo   repository.IComic
	EpisodeRepo repository.IEpisode
	PageRepo    repository.IPage

	Kendang connection.IKendang
}

// Perform ...
func (r Read) Perform(comicName string, episodeNo float64) ([]string, error) {
	comic, err := r.ComicRepo.Find(comicName)
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

	page, err := r.PageRepo.FindByEpisode(episode.ID, sources[n])
	if err != nil {
		return nil, err
	}

	if len(page.Links) < 1 {
		err = r.fetchFromKendang(&page)
	}

	return page.Links, err
}

func (r Read) fetchFromKendang(page *model.Page) error {
	pagesLink, err := r.Kendang.FetchPages(page.Link, page.SourceID.Hex())
	if err != nil {
		return err
	}

	page.Links = pagesLink
	return r.PageRepo.Update(page)
}
