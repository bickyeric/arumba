package resolver

import (
	"context"

	"github.com/bickyeric/arumba/generated"
	"github.com/bickyeric/arumba/model"
)

type comic struct{ generated.ResolverRoot }

func (r *comic) Episodes(ctx context.Context, comic *model.Comic, after *string, first *int) (*model.EpisodeConnection, error) {
	return r.Query().Episodes(ctx, comic.ID, after, first)
}
