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
	var comics []*model.Comic
	res, err := r.cRepo.FindAll(name)
	for i := 0; i < len(res); i++ {
		comics = append(comics, &res[i])
	}
	return comics, err
}
