package comic

import (
	"github.com/bickyeric/arumba/model"
	"github.com/bickyeric/arumba/repository"
)

// Searcher ...
type Searcher interface {
	Perform(name string) ([]model.Comic, error)
}

// NewSearch ...
func NewSearch(repo repository.IComic) Searcher {
	return search{repo}
}

type search struct {
	repo repository.IComic
}

func (s search) Perform(name string) ([]model.Comic, error) {
	return s.repo.FindAll(name)
}
