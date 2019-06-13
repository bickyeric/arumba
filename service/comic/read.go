package comic

import (
	"math/rand"
	"time"

	"github.com/bickyeric/arumba"
	"github.com/bickyeric/arumba/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Read ...
type Read struct {
	SourceRepo  repository.ISource
	ComicRepo   repository.IComic
	EpisodeRepo repository.IEpisode
	PageRepo    repository.IPage
}

// NewRead ...
func NewRead(app arumba.Arumba) Read {
	return Read{
		SourceRepo:  app.SourceRepo,
		ComicRepo:   app.ComicRepo,
		EpisodeRepo: app.EpisodeRepo,
		PageRepo:    app.PageRepo,
	}
}

// PerformByComicName ...
func (r Read) PerformByComicName(comicName string, episodeNo float64) (string, error) {
	comic, err := r.ComicRepo.Find(comicName)
	if err != nil {
		return "", err
	}

	return r.PerformByComicID(comic.ID, episodeNo)
}

// PerformByComicID ...
func (r Read) PerformByComicID(id primitive.ObjectID, episodeNo float64) (string, error) {
	comic, err := r.ComicRepo.FindByID(id)
	if err != nil {
		return "", err
	}

	episode, err := r.EpisodeRepo.FindByNo(comic.ID, episodeNo)
	if err != nil {
		return "", err
	}

	sourceIDs, err := r.PageRepo.GetSources(episode.ID)
	if err != nil {
		return "", err
	}

	rand.Seed(time.Now().Unix())
	n := rand.Int() % len(sourceIDs)

	source, err := r.SourceRepo.FindByID(sourceIDs[n])
	if err != nil {
		return "", err
	}

	page, err := r.PageRepo.FindByEpisode(episode.ID, source.ID)
	return page.TelegraphLink, err
}
