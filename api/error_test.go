package api_test

import (
	"errors"
	"testing"

	"github.com/bickyeric/arumba/api"
	"github.com/stretchr/testify/suite"
)

type httpErrorSuite struct {
	suite.Suite
}

func (s httpErrorSuite) TestResponseHeaders() {
	err := api.NewHTTPError(errors.New("id is not object id"), 400, "provided id is invalid")
	s.NotPanics(func() {
		clientErr := err.(api.ClientError)
		status, header := clientErr.ResponseHeaders()
		s.NotNil(header)
		s.Equal(400, status)
	})
}

func (s httpErrorSuite) TestError() {
	err := api.NewHTTPError(errors.New("Not Found"), 404, "Resource is not found")
	s.NotPanics(func() {
		clientErr := err.(api.ClientError)
		s.Equal("Resource is not found : Not Found", clientErr.Error())
	})
}

func TestHTTPError(t *testing.T) {
	suite.Run(t, new(httpErrorSuite))
}
