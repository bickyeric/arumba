package comic

import (
	"github.com/bickyeric/arumba/model"
	"github.com/bickyeric/arumba/repository"
)

// ISearch ...
type ISearch interface {
	Perform(name string) ([]model.Comic, error)
}

// Search ...
type Search struct {
	Repo repository.IComic
}

// Perform ...
func (s Search) Perform(name string) ([]model.Comic, error) {
	return s.Repo.Search(name)
}
