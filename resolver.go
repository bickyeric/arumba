package arumba

import (
	"github.com/bickyeric/arumba/generated"
	"github.com/bickyeric/arumba/repository"
	"github.com/bickyeric/arumba/resolver"
)

// Resolver ...
type Resolver struct {
	ComicRepo   repository.IComic
	EpisodeRepo repository.IEpisode
	SourceRepo  repository.ISource
}

// Comic ...
func (r *Resolver) Comic() generated.ComicResolver {
	return resolver.NewComic(r.EpisodeRepo)
}

// Episode ...
func (r *Resolver) Episode() generated.EpisodeResolver {
	return resolver.NewEpisode(r.ComicRepo)
}

// Source ...
func (r *Resolver) Source() generated.SourceResolver {
	return resolver.NewSource()
}

// Query ...
func (r *Resolver) Query() generated.QueryResolver {
	return resolver.NewQuery(r.ComicRepo, r.SourceRepo)
}
