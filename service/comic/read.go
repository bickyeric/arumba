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
func (r Read) Perform(comicName string, episodeNo float64) ([]*model.Page, error) {
	comic, err := r.ComicRepo.FindOne(comicName)
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

	pages, err := r.PageRepo.FindByEpisode(episode.ID, sources[n])
	if err != nil {
		return nil, err
	}

	if len(pages) < 1 {
		pages, err = r.fetchFromKendang(episode.ID, sources[n])
	}

	return pages, err
}

func (r Read) fetchFromKendang(episodeID, sourceID int) ([]*model.Page, error) {
	episodeLink, err := r.EpisodeRepo.GetLink(episodeID, sourceID)
	if err != nil {
		return nil, err
	}

	pagesLink, err := r.Kendang.FetchPages(episodeLink, sourceID)
	if err != nil {
		return nil, err
	}

	result := []*model.Page{}
	for _, link := range pagesLink {
		page := model.Page{
			Link:      link,
			EpisodeID: episodeID,
			SourceID:  sourceID,
		}
		err := r.PageRepo.Insert(&page)
		if err != nil {
			return nil, err
		}
		result = append(result, &page)
	}
	return result, nil
}
