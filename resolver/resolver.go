package resolver

import (
	"github.com/bickyeric/arumba/external"
	"github.com/bickyeric/arumba/generated"
	"github.com/bickyeric/arumba/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

// DefaultValue
const (
	DefaultFirst  = 20
	DefaultOffset = 0
)

type root struct {
	comicColl, episodeColl *mongo.Collection
	sourceRepository       repository.ISource
	pageRepository         repository.IPage
}

// New create graphql root resolver
func New(db external.MongoDatabase) generated.ResolverRoot {
	return &root{
		episodeColl:      db.Collection("episodes"),
		comicColl:        db.Collection("comics"),
		sourceRepository: repository.NewSource(db),
		pageRepository:   repository.NewPage(db),
	}
}

func (r *root) Query() generated.QueryResolver {
	return &query{r}
}

func (r *root) Comic() generated.ComicResolver {
	return &comicResolver{r}
}

func (r *root) Episode() generated.EpisodeResolver {
	return &episodeResolver{r, r.pageRepository}
}

func (r *root) EpisodeConnection() generated.EpisodeConnectionResolver {
	return &episodeConnection{r, r.episodeColl}
}

func (r *root) ComicConnection() generated.ComicConnectionResolver {
	return &comicConnection{r, r.comicColl}
}

func (r *root) Mutation() generated.MutationResolver {
	return &mutation{r}
}
