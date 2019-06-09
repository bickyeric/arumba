package comic

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bickyeric/arumba"
	"github.com/bickyeric/arumba/connection"
	"github.com/bickyeric/arumba/model"
	"github.com/bickyeric/arumba/repository"
	"github.com/bickyeric/arumba/service/telegraph"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Read ...
type Read struct {
	SourceRepo  repository.ISource
	ComicRepo   repository.IComic
	EpisodeRepo repository.IEpisode
	PageRepo    repository.IPage

	Kendang     connection.IKendang
	PageCreator telegraph.PageCreator
}

// NewRead ...
func NewRead(app arumba.Arumba, kendang connection.IKendang, pageCreator telegraph.PageCreator) Read {
	return Read{
		SourceRepo:  app.SourceRepo,
		ComicRepo:   app.ComicRepo,
		EpisodeRepo: app.EpisodeRepo,
		PageRepo:    app.PageRepo,
		Kendang:     kendang,
		PageCreator: pageCreator,
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
	if err != nil {
		return "", err
	}

	if page.TelegraphLink != "" {
		return page.TelegraphLink, nil
	}

	err = r.generateTelegraphURL(source, comic, *episode, &page)

	return page.TelegraphLink, err
}

func (r Read) generateTelegraphURL(source model.Source, comic model.Comic, episode model.Episode, page *model.Page) (err error) {
	if len(page.Links) < 1 {
		if err = r.fetchFromKendang(page); err != nil {
			return err
		}
	}

	url, err := r.PageCreator.Perform(source.Name, fmt.Sprintf("%s %.1f | %s", comic.Name, episode.No, episode.Name), page.Links)
	if err != nil {
		return err
	}

	page.TelegraphLink = url
	return r.PageRepo.Update(page)
}

func (r Read) fetchFromKendang(page *model.Page) error {
	pagesLink, err := r.Kendang.FetchPages(page.Link, page.SourceID.Hex())
	if err != nil {
		return err
	}

	page.Links = pagesLink
	return r.PageRepo.Update(page)
}
