package resolver_test

import (
	"context"
	"testing"

	"github.com/bickyeric/arumba/model"
	"github.com/bickyeric/arumba/repository/mock"
	"github.com/bickyeric/arumba/resolver"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type querySuite struct {
	suite.Suite
}

func (s *querySuite) TestComics() {
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()

	q := "one piece"
	first, offset := 20, 0
	expected := model.Comic{
		Name: "One Piece",
	}
	comicRepo := mock.NewMockIComic(ctrl)
	comicRepo.EXPECT().FindAll(q, first, offset).Return([]model.Comic{expected}, nil)
	query := resolver.NewQuery(comicRepo)
	actual, err := query.Comics(context.Background(), q, &first, &offset)
	s.Nil(err)
	for _, comic := range actual {
		s.Equal(expected.Name, comic.Name)
	}
}

func TestQuery(t *testing.T) {
	suite.Run(t, new(querySuite))
}
