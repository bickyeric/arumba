package resolver

import (
	"context"

	"github.com/bickyeric/arumba/generated"
	"github.com/bickyeric/arumba/model"
	"github.com/bickyeric/arumba/repository"
)

type comic struct {
	eRepo repository.IEpisode
}

func NewComic(eRepo repository.IEpisode) generated.ComicResolver {
	return &comic{eRepo: eRepo}
}

func (r *comic) Episodes(ctx context.Context, comic *model.Comic, first, offset *int) ([]*model.Episode, error) {
	var episodes []*model.Episode
	f, o := 100, 0
	if first != nil {
		f = *first
	}
	if offset != nil {
		o = *offset
	}
	res, err := r.eRepo.FindAll(ctx, comic.ID, f, o)
	for i := 0; i < len(res); i++ {
		episodes = append(episodes, &res[i])
	}
	return episodes, err
}
