package episode

import (
	"github.com/bickyeric/arumba/repository"
)

type Search struct {
	Repo repository.IEpisode
}

func (s Search) Perform(comicID int) ([][]float64, error) {
	totalEpisode, err := s.Repo.Count(comicID)
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
		no, err := s.Repo.No(comicID, index)
		if err != nil {
			return nil, err
		}
		noRange = append(noRange, no)

		index += member
		if member > 1 {
			no, err := s.Repo.No(comicID, index-1)
			if err != nil {
				return nil, err
			}
			noRange = append(noRange, no)
		}
		noGroup = append(noGroup, noRange)
	}
	return noGroup, nil
}
