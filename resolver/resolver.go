package resolver

import (
	"github.com/bickyeric/arumba/generated"
	"go.mongodb.org/mongo-driver/mongo"
)

// DefaultValue
const (
	DefaultFirst  = 20
	DefaultOffset = 0
)

type resolver struct {
	episode generated.EpisodeResolver
	db      *mongo.Database
}

// New create graphql root resolver
func New(episode generated.EpisodeResolver, db *mongo.Database) generated.ResolverRoot {
	return &resolver{
		episode: episode,
		db:      db,
	}
}

func (r *resolver) Query() generated.QueryResolver {
	return &query{r}
}

func (r *resolver) Comic() generated.ComicResolver {
	return &comicResolver{r}
}

func (r *resolver) Episode() generated.EpisodeResolver {
	return r.episode
}

func (r *resolver) EpisodeConnection() generated.EpisodeConnectionResolver {
	return &episodeConnection{r, r.db.Collection("episodes")}
}

func (r *resolver) ComicConnection() generated.ComicConnectionResolver {
	return &comicConnection{r, r.db.Collection("comics")}
}
