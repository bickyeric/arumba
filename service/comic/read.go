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
	sourceRepo  repository.ISource
	comicRepo   repository.IComic
	episodeRepo repository.IEpisode
	pageRepo    repository.IPage
}

// NewRead ...
func NewRead(app arumba.Arumba) Read {
	return Read{
		sourceRepo:  app.SourceRepo,
		episodeRepo: app.EpisodeRepo,
		pageRepo:    app.PageRepo,
	}
}

// PerformByComicName ...
func (r Read) PerformByComicName(comicName string, episodeNo float64) (string, error) {
	comic, err := r.comicRepo.Find(comicName)
	if err != nil {
		return "", err
	}

	return r.PerformByComicID(comic.ID, episodeNo)
}

// PerformByComicID ...
func (r Read) PerformByComicID(comicID primitive.ObjectID, episodeNo float64) (string, error) {
	episode, err := r.episodeRepo.FindByNo(comicID, episodeNo)
	if err != nil {
		return "", err
	}

	sourceIDs, err := r.pageRepo.GetSources(episode.ID)
	if err != nil {
		return "", err
	}

	rand.Seed(time.Now().Unix())
	n := rand.Int() % len(sourceIDs)

	source, err := r.sourceRepo.FindByID(sourceIDs[n])
	if err != nil {
		return "", err
	}

	page, err := r.pageRepo.FindByEpisode(episode.ID, source.ID)
	return page.Link, err
}
