package comic

import (
	"github.com/bickyeric/arumba/model"
	"github.com/bickyeric/arumba/repository"
)

// Finder ...
type Finder interface {
	Perform(name string) (model.Comic, error)
}

// NewFinder ...
func NewFinder(repo repository.IComic) Finder {
	return find{repo}
}

type find struct {
	repo repository.IComic
}

func (f find) Perform(name string) (model.Comic, error) {
	return f.repo.Find(name)
}
