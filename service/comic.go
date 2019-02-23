package service

import (
	"math/rand"
	"time"

	"github.com/bickyeric/arumba/model"
	"github.com/bickyeric/arumba/repository"
)

type IComic interface {
	ReadComic(comicName string, episodeNo float64) ([]*model.Page, error)
}

type ComicService struct {
	ComicRepo   repository.IComic
	EpisodeRepo repository.IEpisode
	PageRepo    repository.IPage
}

func (service ComicService) ReadComic(comicName string, episodeNo float64) ([]*model.Page, error) {
	comic, err := service.ComicRepo.FindByName(comicName)
	if err != nil {
		return nil, err
	}

	episode, err := service.EpisodeRepo.FindByNo(comic.ID, episodeNo)
	if err != nil {
		return nil, err
	}

	sources := service.EpisodeRepo.GetSources(episode.ID)

	rand.Seed(time.Now().Unix())
	n := rand.Int() % len(sources)

	return service.PageRepo.FindByEpisode(episode.ID, sources[n])
}
