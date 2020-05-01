package middleware_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bickyeric/arumba/api/middleware"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/suite"
)

type basicAuthSuite struct {
	suite.Suite
}

func (s *basicAuthSuite) TestNotAuthenticated() {
	nextResolver := func(c echo.Context) error {
		return nil
	}
	auth := middleware.BasicAuth{
		Username: "root",
		Password: "root",
	}

	req, _ := http.NewRequest("GET", "http://local.host", nil)
	echoContext := echo.New().NewContext(req, httptest.NewRecorder())
	err := auth.Checker(nextResolver)(echoContext)
	s.Nil(err)

	ctx := echoContext.Request().Context()
	_, err = auth.IsAuthenticated(ctx, nil, nil)
	s.NotNil(err)
}

func (s *basicAuthSuite) TestAuthenticated() {
	next := func(c echo.Context) error {
		return nil
	}
	nextResolver := func(ctx context.Context) (res interface{}, err error) {
		return nil, nil
	}
	auth := middleware.BasicAuth{
		Username: "root",
		Password: "root",
	}

	req, _ := http.NewRequest("GET", "http://local.host", nil)
	req.SetBasicAuth("root", "root")
	echoContext := echo.New().NewContext(req, httptest.NewRecorder())
	err := auth.Checker(next)(echoContext)
	s.Nil(err)

	ctx := echoContext.Request().Context()
	_, err = auth.IsAuthenticated(ctx, nil, nextResolver)
	s.Nil(err)
}

func TestBasicAuth(t *testing.T) {
	suite.Run(t, new(basicAuthSuite))
}
