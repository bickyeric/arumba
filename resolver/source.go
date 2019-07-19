package resolver

import (
	"context"

	"github.com/bickyeric/arumba/generated"
	"github.com/bickyeric/arumba/model"
)

type source struct{}

// NewSource ...
func NewSource() generated.SourceResolver {
	return source{}
}

func (r source) ID(ctx context.Context, obj *model.Source) (string, error) {
	return obj.ID.Hex(), nil
}
