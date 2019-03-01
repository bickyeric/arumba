package episode

import (
	"database/sql"
	"errors"
	"time"

	"github.com/bickyeric/arumba/model"
	"github.com/bickyeric/arumba/repository"
)

var ErrEpisodeExist = errors.New("episode exist")

type UpdateSaver struct {
	ComicRepo   repository.IComic
	EpisodeRepo repository.IEpisode
	PageRepo    repository.IPage
}

func (s UpdateSaver) Perform(update model.Update, sourceID int) error {
	comic, err := s.getComic(update.ComicName)
	if err != nil {
		return err
	}

	episode, err := s.getEpisode(comic.ID, update)
	if err != nil {
		return err
	}

	if _, err := s.EpisodeRepo.GetLink(episode.ID, sourceID); err == nil {
		return ErrEpisodeExist
	}

	return s.EpisodeRepo.InsertLink(episode.ID, sourceID, update.EpisodeLink)
}

func (s UpdateSaver) getComic(name string) (model.Comic, error) {
	var comic, err = s.ComicRepo.FindOne(name)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			comic.Name = name
			comic.CreatedAt = time.Now()
			comic.UpdatedAt = time.Now()
			return comic, s.ComicRepo.Insert(&comic)
		default:
			return comic, err
		}
	}
	return comic, nil
}

func (s UpdateSaver) getEpisode(comicID int, update model.Update) (*model.Episode, error) {
	episode, err := s.EpisodeRepo.FindByNo(comicID, update.EpisodeNo)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			episode := new(model.Episode)
			episode.No = update.EpisodeNo
			episode.Name = update.EpisodeName
			episode.CreatedAt = time.Now()
			episode.UpdatedAt = time.Now()
			episode.ComicID = comicID
			return episode, s.EpisodeRepo.Insert(episode)
		default:
			return nil, err
		}
	}

	return episode, nil
}
