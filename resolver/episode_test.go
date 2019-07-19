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

func TestEpisodeID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	episodeID := primitive.NewObjectID()

	comicRepo := mocks.NewMockIComic(ctrl)
	r := resolver.NewEpisode(comicRepo)
	stringID, err := r.ID(context.Background(), &model.Episode{
		ID: episodeID,
	})

	assert.Nil(t, err)
	assert.Equal(t, episodeID.Hex(), stringID)
}

func TestEpisodesComic(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	comicID := primitive.NewObjectID()
	expectedResult := model.Comic{}

	comicRepo := mocks.NewMockIComic(ctrl)
	comicRepo.EXPECT().FindByID(comicID).Return(expectedResult, nil)

	r := resolver.NewEpisode(comicRepo)
	results, err := r.Comic(context.Background(), &model.Episode{
		ComicID: comicID,
	})

	assert.Nil(t, err)
	assert.Equal(t, &expectedResult, results)
}
