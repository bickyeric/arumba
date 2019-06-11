package episode

import (
	"github.com/bickyeric/arumba/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Searcher ...
type Searcher interface {
	Perform(comicID primitive.ObjectID, bound ...float64) ([][]float64, error)
}

// NewSearch ...
func NewSearch(repo repository.IEpisode) Searcher {
	return search{repo}
}

type search struct {
	repo repository.IEpisode
}

// Perform ...
func (s search) Perform(comicID primitive.ObjectID, bound ...float64) ([][]float64, error) {
	totalEpisode, err := s.repo.Count(comicID, bound...)
	if err != nil {
		return nil, err
	}
	noGroup := [][]float64{}

	number := 5
	index := 0
	for i := 0; i < number; i++ {
		member := (totalEpisode - index) / (number - i)
		if member < 1 {
			continue
		}
		noRange := []float64{}
		no, err := s.repo.No(comicID, index, bound...)
		if err != nil {
			return nil, err
		}
		noRange = append(noRange, no)

		index += member
		if member > 1 {
			no, err := s.repo.No(comicID, index-1, bound...)
			if err != nil {
				return nil, err
			}
			noRange = append(noRange, no)
		}
		noGroup = append(noGroup, noRange)
	}
	return noGroup, nil
}
