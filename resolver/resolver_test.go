package resolver_test

import (
	"testing"

	"github.com/bickyeric/arumba/resolver"
	"github.com/stretchr/testify/suite"
)

type rootResolverSuite struct {
	suite.Suite
}

func (s *rootResolverSuite) TestResolver() {
	resolver := resolver.New(nil)
	s.NotPanics(func() {
		s.Nil(resolver.Query())
	})
}

func TestRootResolver(t *testing.T) {
	suite.Run(t, new(rootResolverSuite))
}
