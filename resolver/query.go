package resolver

import (
	"context"

	"github.com/bickyeric/arumba/generated"
	"github.com/bickyeric/arumba/model"
	"github.com/bickyeric/arumba/repository"
)

type query struct {
	cRepo repository.IComic
}

func NewQuery(cRepo repository.IComic) generated.QueryResolver {
	return &query{cRepo: cRepo}
}

func (r *query) Comics(ctx context.Context, name string) ([]*model.Comic, error) {
	return r.cRepo.FindAll(name)
}
