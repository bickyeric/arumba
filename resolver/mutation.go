package resolver

import (
	"context"

	"github.com/bickyeric/arumba/model"
)

type mutation struct{ *root }

func (m *mutation) SourceCreate(ctx context.Context, input model.SourceInput) (*model.SourceCreatePayload, error) {
	source := model.Source{
		Name:     input.Name,
		Hostname: input.Hostname,
	}
	payload := &model.SourceCreatePayload{
		Source: &source,
	}
	err := m.sourceRepository.Insert(&source)
	if err != nil {
		payload.UserError = append(payload.UserError, &model.UserError{Message: err.Error()})
	}
	return payload, nil
}
