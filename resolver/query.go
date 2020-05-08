package resolver

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/bickyeric/arumba/model"
	"github.com/bickyeric/arumba/resolver/comic"
	"github.com/bickyeric/arumba/resolver/pagination"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type query struct{ *root }

func (r *query) Comics(ctx context.Context, name string, after *string, first *int) (conn *model.ComicConnection, err error) {
	conn = new(model.ComicConnection)
	conn.BaseQuery = primitive.M{"name": primitive.M{"$regex": ".*" + name + ".*", "$options": "i"}}
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

func (r *query) Sources(ctx context.Context) (sources []*model.Source, err error) {
	cur, err := r.sourceColl.Find(ctx, primitive.M{})
	if err != nil {
		return sources, err
	}
	err = cur.All(ctx, &sources)
	return sources, err
}

func (r *query) Source(ctx context.Context, sourceID primitive.ObjectID) (source *model.Source, err error) {
	source = new(model.Source)
	err = r.sourceColl.FindOne(ctx, primitive.M{"_id": sourceID}).Decode(source)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return source, err
}
