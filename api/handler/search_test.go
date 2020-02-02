package handler_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bickyeric/arumba/api"
	"github.com/bickyeric/arumba/api/handler"
	"github.com/bickyeric/arumba/model"
	"github.com/bickyeric/arumba/service/comic/mock"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ctx = context.Background()

type searchSuite struct {
	suite.Suite

	e              *echo.Echo
	ctrl           *gomock.Controller
	mockedSearcher *mock.MockSearcher
	rec            *httptest.ResponseRecorder
}

func (s *searchSuite) SetupTest() {
	s.e = echo.New()
	s.ctrl = gomock.NewController(s.T())
	s.mockedSearcher = mock.NewMockSearcher(s.ctrl)
	s.rec = httptest.NewRecorder()
}

func (s *searchSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *searchSuite) TestOk() {
	expected := []model.Comic{
		model.Comic{
			ID:   primitive.NewObjectID(),
			Name: "One Piece",
		},
	}

	s.mockedSearcher.EXPECT().Perform(ctx, "one piece").Return(expected, nil)
	req := httptest.NewRequest(http.MethodGet, "/search?q=one%20piece", nil)
	req.Header.Set("Content-Type", echo.MIMEApplicationJSON)
	c := s.e.NewContext(req, s.rec)

	handler := handler.NewSearch(s.mockedSearcher)
	s.Nil(handler.OnHandle(c))

	var actual []model.Comic
	result := s.rec.Result()
	json.NewDecoder(result.Body).Decode(&actual)
	s.Equal(expected, actual)
}

func (s *searchSuite) TestBlankQuery() {
	req := httptest.NewRequest(http.MethodGet, "/search", nil)
	req.Header.Set("Content-Type", echo.MIMEApplicationJSON)
	c := s.e.NewContext(req, s.rec)

	handler := handler.NewSearch(s.mockedSearcher)
	err := handler.OnHandle(c)
	s.NotNil(err)
	clientError, ok := err.(*api.HTTPError)
	s.True(ok)
	s.Equal(http.StatusUnprocessableEntity, clientError.Status)
}

func (s *searchSuite) TestUnexpectedError() {
	expectedErr := errors.New("unexpected-error")
	s.mockedSearcher.EXPECT().Perform(ctx, "one piece").Return(nil, expectedErr)
	req := httptest.NewRequest(http.MethodGet, "/search?q=one%20piece", nil)
	req.Header.Set("Content-Type", echo.MIMEApplicationJSON)
	c := s.e.NewContext(req, s.rec)

	handler := handler.NewSearch(s.mockedSearcher)
	err := handler.OnHandle(c)
	s.Equal(err, expectedErr)
}

func TestKendangWebhook(t *testing.T) {
	suite.Run(t, new(searchSuite))
}
