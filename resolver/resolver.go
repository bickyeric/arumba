package resolver

import (
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
	episode                generated.EpisodeResolver
	comicColl, episodeColl *mongo.Collection
	sourceRepository       repository.ISource
}

// New create graphql root resolver
func New(episode generated.EpisodeResolver, db *mongo.Database) generated.ResolverRoot {
	return &root{
		episode:          episode,
		episodeColl:      db.Collection("episodes"),
		comicColl:        db.Collection("comics"),
		sourceRepository: repository.NewSource(db),
	}
}

func (r *root) Query() generated.QueryResolver {
	return &query{r}
}

func (r *root) Comic() generated.ComicResolver {
	return &comicResolver{r}
}

func (r *root) Episode() generated.EpisodeResolver {
	return r.episode
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
