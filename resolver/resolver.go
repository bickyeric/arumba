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
	query   generated.QueryResolver
	comic   generated.ComicResolver
	episode generated.EpisodeResolver
}

func New(q generated.QueryResolver, comic generated.ComicResolver, episode generated.EpisodeResolver) generated.ResolverRoot {
	return &resolver{
		query:   q,
		comic:   comic,
		episode: episode,
	}
}

func (r *resolver) Query() generated.QueryResolver {
	return r.query
}

func (r *resolver) Comic() generated.ComicResolver {
	return r.comic
}

func (r *resolver) Episode() generated.EpisodeResolver {
	return r.episode
}
