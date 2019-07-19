package resolver

import (
	"context"

	"github.com/bickyeric/arumba/generated"
	"github.com/bickyeric/arumba/model"
	"github.com/bickyeric/arumba/repository"
)

type episode struct {
	comicRepo repository.IComic
}

// NewEpisode ...
func NewEpisode(comicRepo repository.IComic) generated.EpisodeResolver {
	return episode{comicRepo}
}

func (r episode) ID(ctx context.Context, obj *model.Episode) (string, error) {
	return obj.ID.Hex(), nil
}

func (r episode) Comic(ctx context.Context, obj *model.Episode) (*model.Comic, error) {
	c, err := r.comicRepo.FindByID(obj.ComicID)
	return &c, err
}
