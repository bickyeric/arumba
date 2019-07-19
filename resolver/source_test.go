package resolver_test

import (
	"context"
	"testing"

	"github.com/bickyeric/arumba/model"
	"github.com/bickyeric/arumba/resolver"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestSourceID(t *testing.T) {
	sourceID := primitive.NewObjectID()

	r := resolver.NewSource()
	stringID, err := r.ID(context.Background(), &model.Source{
		ID: sourceID,
	})

	assert.Nil(t, err)
	assert.Equal(t, sourceID.Hex(), stringID)
}
