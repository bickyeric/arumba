package resolver

import (
	"context"

	"github.com/bickyeric/arumba/resolver/comic"

	"github.com/bickyeric/arumba/generated"
	"github.com/bickyeric/arumba/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type comicResolver struct{ generated.ResolverRoot }

func (r *comicResolver) Episodes(ctx context.Context, comic *model.Comic, after *string, first *int) (*model.EpisodeConnection, error) {
	return r.Query().Episodes(ctx, comic.ID, after, first)
}

type comicConnection struct {
	generated.ResolverRoot
	*mongo.Collection
}

func (r *comicConnection) Edges(ctx context.Context, c *model.ComicConnection) (edges []*model.ComicEdge, err error) {
	edges = make([]*model.ComicEdge, 0)
	opts := options.Find().SetLimit(int64(c.Limit)).SetSkip(int64(c.Skip))
	cur, err := r.Find(ctx, c.BaseQuery, opts)
	if err != nil {
		return edges, err
	}

	var i int = c.Skip + 1
	for cur.Next(ctx) {
		edge := model.ComicEdge{
			Cursor: comic.EncodeCursor(i),
			Node:   new(model.Comic),
		}
		if err := cur.Decode(edge.Node); err != nil {
			return edges, err
		}
		edges = append(edges, &edge)
		i++
	}

	return edges, err
}

func (r *comicConnection) PageInfo(ctx context.Context, c *model.ComicConnection) (pageInfo *model.PageInfo, err error) {
	pageInfo = new(model.PageInfo)
	opts := options.Find().SetLimit(2).SetSkip(int64(c.Skip + c.Limit - 1))
	cur, err := r.Find(ctx, c.BaseQuery, opts)
	if err != nil {
		return pageInfo, err
	}

	var nextInfo []model.Comic
	if err = cur.All(ctx, &nextInfo); err != nil {
		return pageInfo, err
	}
	pageInfo.HasNextPage = len(nextInfo) == 2
	if pageInfo.HasNextPage {
		pageInfo.StartCursor = comic.EncodeCursor(c.Skip + c.Limit)
	}
	return pageInfo, err
}
