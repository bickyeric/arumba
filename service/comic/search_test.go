package comic_test

import (
	"context"
	"testing"

	"github.com/bickyeric/arumba/model"
	"github.com/bickyeric/arumba/repository/mock"
	"github.com/bickyeric/arumba/service/comic"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewSearch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	comicRepo := mock.NewMockIComic(ctrl)
	assert.NotPanics(t, func() {
		comic.NewSearch(comicRepo)
	})
}

func TestPerform(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	comicRepo := mock.NewMockIComic(ctrl)
	comicRepo.EXPECT().FindAll(ctx, "One Piece", 20, 0).Return([]model.Comic{model.Comic{Name: "One Piece"}}, nil)

	searcher := comic.NewSearch(comicRepo)
	comics, err := searcher.Perform(ctx, "One Piece")
	assert.Nil(t, err)
	assert.Len(t, comics, 1)
}
