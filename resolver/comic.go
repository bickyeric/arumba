package resolver

import (
	"context"

	"github.com/bickyeric/arumba/generated"
	"github.com/bickyeric/arumba/model"
	"github.com/bickyeric/arumba/repository"
)

type comic struct {
	episodeRepo repository.IEpisode
}

// NewComic ...
func NewComic(repo repository.IEpisode) generated.ComicResolver {
	return comic{repo}
}

func (r comic) ID(ctx context.Context, obj *model.Comic) (string, error) {
	return obj.ID.Hex(), nil
}

func (r comic) Episodes(ctx context.Context, obj *model.Comic) ([]*model.Episode, error) {
	return r.episodeRepo.AllByComicID(obj.ID)
}
