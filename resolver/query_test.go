package resolver_test

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/bickyeric/arumba/model"
	"github.com/bickyeric/arumba/repository/mock"
	"github.com/bickyeric/arumba/resolver"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type querySuite struct {
	suite.Suite
	ctx             context.Context
	ctrl            *gomock.Controller
	mockedComicRepo *mock.MockIComic
}

func (s *querySuite) SetupTest() {
	s.ctx = context.Background()
	s.ctrl = gomock.NewController(s.T())
	s.mockedComicRepo = mock.NewMockIComic(s.ctrl)
}

func (s *querySuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *querySuite) TestComics() {
	q := "one piece"
	first, offset := 20, 0
	expected := model.Comic{
		Name: "One Piece",
	}
	s.mockedComicRepo.EXPECT().FindAll(s.ctx, q, first, offset).Return([]model.Comic{expected}, nil)
	query := resolver.NewQuery(s.mockedComicRepo)
	actual, err := query.Comics(s.ctx, q, &first, &offset)
	s.Nil(err)
	for _, comic := range actual {
		s.Equal(expected.Name, comic.Name)
	}
}

func (s *querySuite) TestEpisodes_NotValidParams() {
	comicID := primitive.NewObjectID()
	first := -90
	query := resolver.NewQuery(s.mockedComicRepo)
	_, err := query.Episodes(s.ctx, comicID, nil, nil, &first, nil)
	s.NotNil(err)
}

func (s *querySuite) TestEpisodes_ForwardPagination() {
	comicID := primitive.NewObjectID()
	query := resolver.NewQuery(s.mockedComicRepo)
	connection, err := query.Episodes(s.ctx, comicID, nil, nil, nil, nil)
	s.Nil(err)
	s.Equal(comicID, connection.ComicID)
	s.NotNil(connection.Pagination.Pipelines())
}

func TestQuery(t *testing.T) {
	suite.Run(t, new(querySuite))
}
