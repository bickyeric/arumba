package comic

import (
	"math/rand"
	"time"

	"github.com/bickyeric/arumba/connection"
	"github.com/bickyeric/arumba/model"
	"github.com/bickyeric/arumba/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Read ...
type Read struct {
	ComicRepo   repository.IComic
	EpisodeRepo repository.IEpisode
	PageRepo    repository.IPage

	Kendang connection.IKendang
}

// PerformByComicName ...
func (r Read) PerformByComicName(comicName string, episodeNo float64) ([]string, error) {
	comic, err := r.ComicRepo.Find(comicName)
	if err != nil {
		return nil, err
	}

	return r.PerformByComicID(comic.ID, episodeNo)
}

// PerformByComicID ...
func (r Read) PerformByComicID(comicID primitive.ObjectID, episodeNo float64) ([]string, error) {
	episode, err := r.EpisodeRepo.FindByNo(comicID, episodeNo)
	if err != nil {
		return nil, err
	}

	sources, err := r.PageRepo.GetSources(episode.ID)
	if err != nil {
		return nil, err
	}

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
