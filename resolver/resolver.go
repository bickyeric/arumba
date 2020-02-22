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
	query generated.QueryResolver
	comic generated.ComicResolver
}

func New(q generated.QueryResolver, comic generated.ComicResolver) generated.ResolverRoot {
	return &resolver{
		query: q,
		comic: comic,
	}
}

func (r *resolver) Query() generated.QueryResolver {
	return r.query
}

func (r *resolver) Comic() generated.ComicResolver {
	return r.comic
}
