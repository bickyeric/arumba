package resolver_test

import (
	"testing"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/golang/mock/gomock"

	"github.com/bickyeric/arumba/external/mock"
	"github.com/bickyeric/arumba/resolver"
	"github.com/stretchr/testify/suite"
)

type rootResolverSuite struct {
	suite.Suite
}

func (s *rootResolverSuite) TestResolver() {
	ctrl := gomock.NewController(s.T())

	mockedDB := mock.NewMockMongoDatabase(ctrl)
	mockedDB.EXPECT().Collection(gomock.Any()).Return(&mongo.Collection{}).AnyTimes()

	resolver := resolver.New(mockedDB)
	s.NotPanics(func() {
		s.NotNil(resolver.Query())
		s.NotNil(resolver.Comic())
		s.NotNil(resolver.Episode())
		s.NotNil(resolver.EpisodeConnection())
		s.NotNil(resolver.ComicConnection())
		s.NotNil(resolver.Mutation())
	})
}

func TestRootResolver(t *testing.T) {
	suite.Run(t, new(rootResolverSuite))
}
