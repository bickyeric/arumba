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

func (r *query) Comics(ctx context.Context, name string, first, offset *int) ([]*model.Comic, error) {
	var comics []*model.Comic
	f, o := DefaultFirst, DefaultOffset
	if first != nil {
		f = *first
	}
	if offset != nil {
		o = *offset
	}
	res, err := r.cRepo.FindAll(name, f, o)
	for i := 0; i < len(res); i++ {
		comics = append(comics, &res[i])
	}
	return comics, err
}
