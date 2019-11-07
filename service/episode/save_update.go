package episode

//go:generate mockgen -destination mock/save_update.go -package=mock -source save_update.go

import (
	"errors"

	"github.com/bickyeric/arumba/model"
	"github.com/bickyeric/arumba/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ErrEpisodeExists ...
var ErrEpisodeExists = errors.New("episode exists")

// UpdateSaver ...
type UpdateSaver interface {
	Perform(update model.Update, sourceID primitive.ObjectID) (page model.Page, err error)
}

type saveUpdate struct {
	SourceRepo  repository.ISource
	ComicRepo   repository.IComic
	EpisodeRepo repository.IEpisode
	PageRepo    repository.IPage
}

// NewSaveUpdate ...
func NewSaveUpdate(sourceRepo repository.ISource,
	comicRepo repository.IComic,
	episodeRepo repository.IEpisode,
	pageRepo repository.IPage) UpdateSaver {
	return saveUpdate{
		SourceRepo:  sourceRepo,
		ComicRepo:   comicRepo,
		EpisodeRepo: episodeRepo,
		PageRepo:    pageRepo,
	}
}

func (s saveUpdate) Perform(update model.Update, sourceID primitive.ObjectID) (page model.Page, err error) {
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
		return existing, ErrEpisodeExists
	}

	page.Link = update.EpisodeLink

	return page, s.PageRepo.Insert(&page)
}

func (s saveUpdate) getComic(name string) (model.Comic, error) {
	comic, err := s.ComicRepo.Find(name)
	if err == mongo.ErrNoDocuments {
		comic.Name = name
		err = s.ComicRepo.Insert(&comic)
	}

	return comic, err
}

func (s saveUpdate) getEpisode(comicID primitive.ObjectID, update model.Update) (*model.Episode, error) {
	episode, err := s.EpisodeRepo.FindByNo(comicID, update.EpisodeNo)
	if err == mongo.ErrNoDocuments {
		episode := new(model.Episode)
		episode.No = update.EpisodeNo
		episode.Name = update.EpisodeName
		episode.ComicID = comicID
		err = s.EpisodeRepo.Insert(episode)
	}

	return episode, err
}
