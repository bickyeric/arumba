package resolver

import (
	"context"

	"github.com/bickyeric/arumba/generated"
	"github.com/bickyeric/arumba/model"
	"github.com/bickyeric/arumba/resolver/comic"
	"github.com/bickyeric/arumba/resolver/pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type query struct{ generated.ResolverRoot }

func (r *query) Comics(ctx context.Context, name string, after *string, first *int) (conn *model.ComicConnection, err error) {
	conn = new(model.ComicConnection)
	conn.BaseQuery = bson.M{"name": bson.M{"$regex": ".*" + name + ".*", "$options": "i"}}
	conn.Limit = 10
	if first != nil {
		conn.Limit = *first
	}
	conn.Skip, err = comic.DecodeCursor(after)
	return conn, err
}

func (r *query) Episodes(ctx context.Context, comicID primitive.ObjectID, after *string, first *int) (conn *model.EpisodeConnection, err error) {
	conn = new(model.EpisodeConnection)
	conn.ComicID = comicID
	conn.Pagination, err = pagination.Validate(after, first)
	return conn, err
}
