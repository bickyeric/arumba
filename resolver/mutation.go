package resolver

import (
	"context"

	"github.com/bickyeric/arumba/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (m *mutation) SourceDelete(ctx context.Context, sourceID primitive.ObjectID) (*model.SourceDeletePayload, error) {
	source, err := m.Query().Source(ctx, sourceID)
	if err != nil {
		return nil, err
	}
	payload := &model.SourceDeletePayload{
		Source: source,
	}
	_, err = m.sourceColl.DeleteOne(ctx, primitive.M{"_id": sourceID})
	if err != nil {
		payload.UserError = append(payload.UserError, &model.UserError{Message: err.Error()})
	}
	return payload, nil
}
