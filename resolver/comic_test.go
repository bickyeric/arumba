package resolver_test

import (
	"context"
	"testing"

	"github.com/bickyeric/arumba/mocks"
	"github.com/bickyeric/arumba/model"
	"github.com/bickyeric/arumba/resolver"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestComicID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	comicID := primitive.NewObjectID()

	episodeRepo := mocks.NewMockIEpisode(ctrl)
	r := resolver.NewComic(episodeRepo)
	stringID, err := r.ID(context.Background(), &model.Comic{
		ID: comicID,
	})

	assert.Nil(t, err)
	assert.Equal(t, comicID.Hex(), stringID)
}

func TestComicEpisodes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	comicID := primitive.NewObjectID()
	expectedResult := []*model.Episode{}

	episodeRepo := mocks.NewMockIEpisode(ctrl)
	episodeRepo.EXPECT().AllByComicID(comicID).Return(expectedResult, nil)

	r := resolver.NewComic(episodeRepo)
	results, err := r.Episodes(context.Background(), &model.Comic{
		ID: comicID,
	})

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, results)
}
