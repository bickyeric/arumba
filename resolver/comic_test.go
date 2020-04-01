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

type comicSuite struct {
	suite.Suite
	ctrl              *gomock.Controller
	mockedEpisodeRepo *mock.MockIEpisode
}

func (s *comicSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.mockedEpisodeRepo = mock.NewMockIEpisode(s.ctrl)
}

func (s *comicSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *comicSuite) TestEpisodes() {
	ctx := context.Background()
	first, offset := 100, 0
	comic := model.Comic{
		ID: primitive.NewObjectID(),
	}
	expected := model.Episode{
		Name: "One Piece Chapter 890",
	}
	s.mockedEpisodeRepo.EXPECT().FindAll(ctx, comic.ID, first, offset).Return([]model.Episode{expected}, nil)
	r := resolver.NewComic(s.mockedEpisodeRepo)
	actual, err := r.Episodes(ctx, &comic, &first, &offset)
	s.Nil(err)
	for _, ep := range actual {
		s.Equal(expected.Name, ep.Name)
	}
}

func TestComic(t *testing.T) {
	suite.Run(t, new(comicSuite))
}
