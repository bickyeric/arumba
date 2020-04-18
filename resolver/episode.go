package resolver

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/bickyeric/arumba/generated"
	"github.com/bickyeric/arumba/model"
	"github.com/bickyeric/arumba/repository"
	eResolver "github.com/bickyeric/arumba/resolver/episode"
)

type episode struct {
	pRepo repository.IPage
}

// NewEpisode ...
func NewEpisode(pRepo repository.IPage) generated.EpisodeResolver {
	return &episode{pRepo: pRepo}
}

func (r *episode) Pages(ctx context.Context, episode *model.Episode) ([]*model.Page, error) {
	var pages []*model.Page
	f, o := 5, 0
	res, err := r.pRepo.FindByEpisode(ctx, episode.ID, f, o)
	for i := 0; i < len(res); i++ {
		pages = append(pages, &res[i])
	}
	return pages, err
}

type episodeConnection struct {
	generated.ResolverRoot
	*mongo.Collection
}

func (r *episodeConnection) Edges(ctx context.Context, c *model.EpisodeConnection) (edges []*model.EpisodeEdge, err error) {
	edges = make([]*model.EpisodeEdge, 0)
	pipe := mongo.Pipeline{
		{{Key: "$match", Value: primitive.M{"comic_id": c.ComicID}}},
	}
	pipe = append(pipe, c.Pagination.Pipelines()...)

	cur, err := r.Aggregate(ctx, pipe)
	if err != nil {
		return edges, err
	}
	for cur.Next(ctx) {
		edge := model.EpisodeEdge{
			Node: new(model.Episode),
		}
		if err = cur.Decode(edge.Node); err != nil {
			return edges, err
		}
		edge.Cursor = eResolver.EncodeCursor(edge.Node.No)
		edges = append(edges, &edge)
	}
	return edges, err
}

func (r *episodeConnection) PageInfo(ctx context.Context, c *model.EpisodeConnection) (pageInfo *model.PageInfo, err error) {
	pageInfo = new(model.PageInfo)
	pipe := mongo.Pipeline{
		{{Key: "$match", Value: primitive.M{"comic_id": c.ComicID}}},
	}
	pipe = append(pipe, c.Pagination.NextPipelines()...)
	cur, err := r.Aggregate(ctx, pipe)
	if err != nil {
		return pageInfo, err
	}

	var nodes []model.Episode
	if err = cur.All(ctx, &nodes); err != nil {
		return pageInfo, err
	}
	pageInfo.HasNextPage = len(nodes) > 1
	if len(nodes) > 0 {
		pageInfo.StartCursor = eResolver.EncodeCursor(nodes[0].No)
	}
	return pageInfo, err
}
