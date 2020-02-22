package resolver

import (
	"context"

	"github.com/bickyeric/arumba/generated"
	"github.com/bickyeric/arumba/model"
	"github.com/bickyeric/arumba/repository"
)

type episode struct {
	pRepo repository.IPage
}

func NewEpisode(pRepo repository.IPage) generated.EpisodeResolver {
	return &episode{pRepo: pRepo}
}

func (r *episode) Pages(ctx context.Context, episode *model.Episode) ([]*model.Page, error) {
	var pages []*model.Page
	f, o := 5, 0
	res, err := r.pRepo.FindByEpisode(ctx, episode.ID, f, o)
	for i := 0; i < len(res); i++ {
		pages = append(pages, &res[i])
	}
	return pages, err
}
