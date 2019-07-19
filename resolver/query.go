package resolver

import (
	"context"

	"github.com/bickyeric/arumba/generated"
	"github.com/bickyeric/arumba/model"
	"github.com/bickyeric/arumba/repository"
)

type query struct {
	comicRepo  repository.IComic
	sourceRepo repository.ISource
}

// NewQuery ...
func NewQuery(comicRepo repository.IComic, sourceRepo repository.ISource) generated.QueryResolver {
	return query{comicRepo, sourceRepo}
}

func (r query) Comics(ctx context.Context, skip, limit int) ([]*model.Comic, error) {
	return r.comicRepo.All()
}

func (r query) Sources(ctx context.Context, skip, limit int) ([]*model.Source, error) {
	return r.sourceRepo.All()
}
