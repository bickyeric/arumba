package episode

import (
	"errors"

	"github.com/bickyeric/arumba/model"
	"github.com/bickyeric/arumba/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var ErrEpisodeExists = errors.New("episode exists")

// UpdateSaver ...
type UpdateSaver struct {
	ComicRepo   repository.IComic
	EpisodeRepo repository.IEpisode
	PageRepo    repository.IPage
}

// Perform ...
func (s UpdateSaver) Perform(update model.Update, sourceID primitive.ObjectID) error {
	comic, err := s.getComic(update.ComicName)
	if err != nil {
		return err
	}

	ep, err := s.getEpisode(comic.ID, update)
	if err != nil {
		return err
	}

	page, err := s.PageRepo.FindByEpisode(ep.ID, sourceID)
	if err == nil {
		return err
	}

	page = model.Page{
		EpisodeID: ep.ID,
		SourceID:  sourceID,
		Link:      update.EpisodeLink,
	}

	return s.PageRepo.Insert(&page)
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
