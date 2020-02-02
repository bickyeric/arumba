package comic

//go:generate mockgen -destination mock/search.go -package=mock -source search.go

import (
	"context"

	"github.com/bickyeric/arumba/model"
	"github.com/bickyeric/arumba/repository"
)

// Searcher ...
type Searcher interface {
	Perform(ctx context.Context, name string) ([]model.Comic, error)
}

// NewSearch ...
func NewSearch(repo repository.IComic) Searcher {
	return search{repo}
}

type search struct {
	repo repository.IComic
}

func (s search) Perform(ctx context.Context, name string) ([]model.Comic, error) {
	return s.repo.FindAll(ctx, name, 20, 0)
}
