package resolver

import (
	"context"

	"github.com/bickyeric/arumba/generated"
	"github.com/bickyeric/arumba/model"
	"github.com/bickyeric/arumba/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type query struct {
	cRepo repository.IComic
	eRepo repository.IEpisode
}

func NewQuery(cRepo repository.IComic, eRepo repository.IEpisode) generated.QueryResolver {
	return &query{cRepo: cRepo, eRepo: eRepo}
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
	res, err := r.cRepo.FindAll(ctx, name, f, o)
	for i := 0; i < len(res); i++ {
		comics = append(comics, &res[i])
	}
	return comics, err
}

func (r *query) Episodes(ctx context.Context, comicID primitive.ObjectID, first, offset *int) ([]*model.Episode, error) {
	var episodes []*model.Episode
	f, o := 100, 0
	if first != nil {
		f = *first
	}
	if offset != nil {
		o = *offset
	}
	res, err := r.eRepo.FindAll(ctx, comicID, f, o)
	for i := 0; i < len(res); i++ {
		episodes = append(episodes, &res[i])
	}
	return episodes, err
}
