package resolver_test

import (
	"context"
	"testing"

	"github.com/bickyeric/arumba/mocks"
	"github.com/bickyeric/arumba/model"
	"github.com/bickyeric/arumba/resolver"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestQueryComics(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectedResult := []*model.Comic{}

	comicRepo := mocks.NewMockIComic(ctrl)
	comicRepo.EXPECT().All().Return(expectedResult, nil)

	sourceRepo := mocks.NewMockISource(ctrl)
	r := resolver.NewQuery(comicRepo, sourceRepo)

	comics, err := r.Comics(context.Background(), 1, 1)

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, comics)
}

func TestQuerySources(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectedResult := []*model.Source{}

	comicRepo := mocks.NewMockIComic(ctrl)
	sourceRepo := mocks.NewMockISource(ctrl)
	sourceRepo.EXPECT().All().Return(expectedResult, nil)

	r := resolver.NewQuery(comicRepo, sourceRepo)

	sources, err := r.Sources(context.Background(), 1, 1)

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, sources)
}
