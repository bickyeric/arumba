package comic

import (
	"github.com/bickyeric/arumba/model"
	"github.com/bickyeric/arumba/repository"
)

type ISearch interface {
	Perform(name string) ([]model.Comic, error)
}

type Search struct {
	Repo repository.IComic
}

func (s Search) Perform(name string) ([]model.Comic, error) {
	return s.Repo.Search(name)
}
