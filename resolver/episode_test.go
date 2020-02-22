package resolver_test

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/bickyeric/arumba/model"
	"github.com/bickyeric/arumba/repository/mock"
	"github.com/bickyeric/arumba/resolver"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type episodeSuite struct {
	suite.Suite
	ctrl           *gomock.Controller
	mockedPageRepo *mock.MockIPage
}

func (s *episodeSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.mockedPageRepo = mock.NewMockIPage(s.ctrl)
}

func (s *episodeSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *episodeSuite) TestPages() {
	first, offset := 5, 0
	episode := model.Episode{
		ID:   primitive.NewObjectID(),
		Name: "One Piece Chapter 890",
	}
	page := model.Page{
		EpisodeID: episode.ID,
		Link:      "www.local.host/ep/890",
	}
	s.mockedPageRepo.EXPECT().FindByEpisode(ctx, episode.ID, first, offset).Return([]model.Page{page}, nil)
	r := resolver.NewEpisode(s.mockedPageRepo)
	actual, err := r.Pages(ctx, &episode)
	s.Nil(err)
	for _, p := range actual {
		s.Equal(page.Link, p.Link)
	}
}

func TestEpisode(t *testing.T) {
	suite.Run(t, new(episodeSuite))
}
