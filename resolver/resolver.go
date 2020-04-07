package resolver

import (
	"github.com/bickyeric/arumba/generated"
)

// DefaultValue
const (
	DefaultFirst  = 20
	DefaultOffset = 0
)

type resolver struct {
	query             generated.QueryResolver
	episode           generated.EpisodeResolver
	episodeConnection generated.EpisodeConnectionResolver
}

// New create graphql root resolver
func New(q generated.QueryResolver, episode generated.EpisodeResolver, episodeConnection generated.EpisodeConnectionResolver) generated.ResolverRoot {
	return &resolver{
		query:             q,
		episode:           episode,
		episodeConnection: episodeConnection,
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
	return r.episodeConnection
}
