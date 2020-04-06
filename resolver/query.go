package resolver

import (
	"context"

	"github.com/bickyeric/arumba/generated"
	"github.com/bickyeric/arumba/model"
	"github.com/bickyeric/arumba/repository"
	"github.com/bickyeric/arumba/resolver/pagination"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type query struct {
	cRepo repository.IComic
}

// NewQuery ...
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
	res, err := r.cRepo.FindAll(ctx, name, f, o)
	for i := 0; i < len(res); i++ {
		comics = append(comics, &res[i])
	}
	return comics, err
}

func (r *query) Episodes(ctx context.Context, comicID primitive.ObjectID, after *string, first *int) (conn *model.EpisodeConnection, err error) {
	conn = new(model.EpisodeConnection)
	conn.ComicID = comicID
	conn.Pagination, err = pagination.Validate(after, first)
	return conn, err
}
