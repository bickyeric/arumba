package resolver_test

import (
	"context"
	"testing"

	"github.com/bickyeric/arumba/model"
	"github.com/bickyeric/arumba/repository/mock"
	"github.com/bickyeric/arumba/resolver"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ctx = context.Background()

type querySuite struct {
	suite.Suite
	ctrl              *gomock.Controller
	mockedComicRepo   *mock.MockIComic
	mockedEpisodeRepo *mock.MockIEpisode
}

func (s *querySuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.mockedComicRepo = mock.NewMockIComic(s.ctrl)
	s.mockedEpisodeRepo = mock.NewMockIEpisode(s.ctrl)
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
	s.mockedComicRepo.EXPECT().FindAll(ctx, q, first, offset).Return([]model.Comic{expected}, nil)
	query := resolver.NewQuery(s.mockedComicRepo, s.mockedEpisodeRepo)
	actual, err := query.Comics(ctx, q, &first, &offset)
	s.Nil(err)
	for _, comic := range actual {
		s.Equal(expected.Name, comic.Name)
	}
}

func (s *querySuite) TestEpisodes() {
	first, offset := 100, 0
	expected := model.Episode{
		ComicID: primitive.NewObjectID(),
		Name:    "One Piece Chapter 890",
	}
	s.mockedEpisodeRepo.EXPECT().FindAll(ctx, expected.ComicID, first, offset).Return([]model.Episode{expected}, nil)
	query := resolver.NewQuery(s.mockedComicRepo, s.mockedEpisodeRepo)
	actual, err := query.Episodes(ctx, expected.ComicID, &first, &offset)
	s.Nil(err)
	for _, ep := range actual {
		s.Equal(expected.Name, ep.Name)
	}
}

func TestQuery(t *testing.T) {
	suite.Run(t, new(querySuite))
}
