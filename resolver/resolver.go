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
	query   generated.QueryResolver
	episode generated.EpisodeResolver
	db      *mongo.Database
}

// New create graphql root resolver
func New(q generated.QueryResolver, episode generated.EpisodeResolver, db *mongo.Database) generated.ResolverRoot {
	return &resolver{
		query:   q,
		episode: episode,
		db:      db,
	}
}

func (r *resolver) Query() generated.QueryResolver {
	return r.query
}

func (r *resolver) Comic() generated.ComicResolver {
	return &comic{r}
}

func (r *resolver) Episode() generated.EpisodeResolver {
	return r.episode
}

func (r *resolver) EpisodeConnection() generated.EpisodeConnectionResolver {
	return &episodeConnection{r, r.db.Collection("episodes")}
}
