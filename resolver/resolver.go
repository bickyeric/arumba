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
}

func New(q generated.QueryResolver) generated.ResolverRoot {
	return &resolver{
		query: q,
	}
}

func (r *resolver) Query() generated.QueryResolver {
	return r.query
}
