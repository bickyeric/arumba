package arumba

import (
	"context"

	"github.com/bickyeric/arumba/generated"
	"github.com/bickyeric/arumba/model"
	"github.com/bickyeric/arumba/repository"
)

type Resolver struct {
	ComicRepo   repository.IComic
	EpisodeRepo repository.IEpisode
	SourceRepo  repository.ISource
}

func (r *Resolver) Comic() generated.ComicResolver {
	return &comicResolver{r}
}

func (r *Resolver) Episode() generated.EpisodeResolver {
	return &episodeResolver{r}
}

func (r *Resolver) Source() generated.SourceResolver {
	return &sourceResolver{r}
}

func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

type episodeResolver struct{ *Resolver }

func (r *episodeResolver) ID(ctx context.Context, obj *model.Episode) (string, error) {
	return obj.ID.Hex(), nil
}

func (r *episodeResolver) Comic(ctx context.Context, obj *model.Episode) (*model.Comic, error) {
	c, err := r.ComicRepo.FindByID(obj.ComicID)
	return &c, err
}

type comicResolver struct{ *Resolver }

func (r *comicResolver) ID(ctx context.Context, obj *model.Comic) (string, error) {
	return obj.ID.Hex(), nil
}

func (r *comicResolver) Episodes(ctx context.Context, obj *model.Comic) ([]*model.Episode, error) {
	return r.EpisodeRepo.AllByComicID(obj.ID)
}

type sourceResolver struct{ *Resolver }

func (r *sourceResolver) ID(ctx context.Context, obj *model.Source) (string, error) {
	return obj.ID.Hex(), nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Comics(ctx context.Context, skip, limit int) ([]*model.Comic, error) {
	return r.ComicRepo.All()
}

func (r *queryResolver) Sources(ctx context.Context, skip, limit int) ([]*model.Source, error) {
	return r.SourceRepo.All()
}
