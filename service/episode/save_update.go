package episode

import (
	"errors"
	"fmt"

	"github.com/bickyeric/arumba"
	"github.com/bickyeric/arumba/connection"
	"github.com/bickyeric/arumba/model"
	"github.com/bickyeric/arumba/repository"
	"github.com/bickyeric/arumba/service/telegraph"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ErrEpisodeExists ...
var ErrEpisodeExists = errors.New("episode exists")

// UpdateSaver ...
type UpdateSaver struct {
	SourceRepo  repository.ISource
	ComicRepo   repository.IComic
	EpisodeRepo repository.IEpisode
	PageRepo    repository.IPage

	Kendang    connection.IKendang
	CreatePage telegraph.CreatePage
}

// NewSaveUpdate ...
func NewSaveUpdate(app arumba.Arumba, kendang connection.IKendang, pageCreator telegraph.CreatePage) UpdateSaver {
	return UpdateSaver{
		SourceRepo:  app.SourceRepo,
		ComicRepo:   app.ComicRepo,
		EpisodeRepo: app.EpisodeRepo,
		PageRepo:    app.PageRepo,
		Kendang:     kendang,
		CreatePage:  pageCreator,
	}
}

// Perform ...
func (s UpdateSaver) Perform(update model.Update, sourceID primitive.ObjectID) (page model.Page, err error) {
	source, err := s.SourceRepo.FindByID(sourceID)
	if err != nil {
		return page, err
	}
	page.SourceID = source.ID

	comic, err := s.getComic(update.ComicName)
	if err != nil {
		return page, err
	}

	ep, err := s.getEpisode(comic.ID, update)
	if err != nil {
		return page, err
	}
	page.EpisodeID = ep.ID

	existing, err := s.PageRepo.FindByEpisode(ep.ID, source.ID)
	if err == nil {
		return existing, nil
	}

	page.Link = update.EpisodeLink

	if err := s.fetchFromKendang(&page); err != nil {
		return page, err
	}

	if err := s.generateTelegraphPage(source, comic, *ep, &page); err != nil {
		return page, err
	}

	// return page, s.PageRepo.Insert(&page)
	return page, nil
}

func (s UpdateSaver) generateTelegraphPage(source model.Source, comic model.Comic, episode model.Episode, page *model.Page) error {
	url, err := s.CreatePage.Perform(source.Name, fmt.Sprintf("%s %.1f | %s", comic.Name, episode.No, episode.Name), page.Links)
	if err != nil {
		return err
	}
	page.TelegraphLink = url
	return nil
}

func (s UpdateSaver) fetchFromKendang(page *model.Page) error {
	pagesLink, err := s.Kendang.FetchPages(page.Link, page.SourceID.Hex())
	if err != nil {
		return err
	}

	page.Links = pagesLink
	return nil
}

func (s UpdateSaver) getComic(name string) (model.Comic, error) {
	comic, err := s.ComicRepo.Find(name)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			comic.Name = name
			return comic, s.ComicRepo.Insert(&comic)
		default:
			return comic, err
		}
	}
	return comic, nil
}

func (s UpdateSaver) getEpisode(comicID primitive.ObjectID, update model.Update) (*model.Episode, error) {
	episode, err := s.EpisodeRepo.FindByNo(comicID, update.EpisodeNo)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			episode := new(model.Episode)
			episode.No = update.EpisodeNo
			episode.Name = update.EpisodeName
			episode.ComicID = comicID
			return episode, s.EpisodeRepo.Insert(episode)
		default:
			return nil, err
		}
	}

	return episode, nil
}
