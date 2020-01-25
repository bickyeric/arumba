package resolver

import (
	"github.com/bickyeric/arumba/generated"
)

type Resolver struct {
	query generated.QueryResolver
}

func New(q generated.QueryResolver) generated.ResolverRoot {
	return &Resolver{
		query: q,
	}
}

func (r *Resolver) Query() generated.QueryResolver {
	return r.query
}
