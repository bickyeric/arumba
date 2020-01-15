package middleware_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bickyeric/arumba/api"
	"github.com/bickyeric/arumba/api/middleware"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/suite"
)

type errorHandlerSuite struct {
	suite.Suite
}

func (s errorHandlerSuite) TestOk() {
	okHandler := func(c echo.Context) error {
		return c.String(http.StatusCreated, "OK")
	}
	fn := middleware.ErrorHandler(okHandler)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	err := fn(echo.New().NewContext(req, rec))
	s.Nil(err)

	res := rec.Result()
	s.Equal(http.StatusCreated, res.StatusCode)
}

func (s errorHandlerSuite) TestUnexpectedError() {
	dummyError := errors.New("Unexpected Error")
	okHandler := func(c echo.Context) error {
		return dummyError
	}
	fn := middleware.ErrorHandler(okHandler)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	err := fn(echo.New().NewContext(req, rec))
	s.Equal(dummyError, err)
}

func (s errorHandlerSuite) TestExpectedError() {
	dummyError := errors.New("Not Found")
	okHandler := func(c echo.Context) error {
		return api.NewHTTPError(dummyError, http.StatusNotFound, "")
	}
	fn := middleware.ErrorHandler(okHandler)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	err := fn(echo.New().NewContext(req, rec))
	s.Nil(err)

	res := rec.Result()
	s.Equal(http.StatusNotFound, res.StatusCode)
}

func TestErrorHandler(t *testing.T) {
	suite.Run(t, new(errorHandlerSuite))
}
