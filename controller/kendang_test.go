package controller_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bickyeric/arumba/controller"
	"github.com/bickyeric/arumba/model"
	"github.com/bickyeric/arumba/service/episode"
	"github.com/bickyeric/arumba/service/episode/mock"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type kendangWebhookSuite struct {
	suite.Suite

	e *echo.Echo
}

func (s *kendangWebhookSuite) SetupTest() {
	s.e = echo.New()
}

func (s *kendangWebhookSuite) TestOk() {
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()

	mockedSvc := mock.NewMockUpdateSaver(ctrl)
	mockedSvc.EXPECT().Perform(gomock.Any(), gomock.Any())

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"sourceID":"5d0f6dfbe4e1f617cbbe18b6"}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := s.e.NewContext(req, rec)

	handler := controller.NewKendang(mockedSvc)

	err := handler.OnHandle(c)
	s.Nil(err)
	s.Equal(http.StatusCreated, rec.Result().StatusCode)
}

func (s *kendangWebhookSuite) TestInvalidJSONBody() {
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{"))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := s.e.NewContext(req, rec)

	handler := controller.NewKendang(nil)

	err := handler.OnHandle(c)
	s.NotNil(err)
}

func (s *kendangWebhookSuite) TestInvalidSourceID() {
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"sourceID":"hehe"}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := s.e.NewContext(req, rec)

	handler := controller.NewKendang(nil)

	err := handler.OnHandle(c)
	s.NotNil(err)
}

func (s *kendangWebhookSuite) TestEpisodeExists() {
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()

	mockedSvc := mock.NewMockUpdateSaver(ctrl)
	mockedSvc.EXPECT().Perform(gomock.Any(), gomock.Any()).Return(model.Page{}, episode.ErrEpisodeExists)

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"sourceID":"5d0f6dfbe4e1f617cbbe18b6"}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := s.e.NewContext(req, rec)

	handler := controller.NewKendang(mockedSvc)

	err := handler.OnHandle(c)
	s.Nil(err)
	s.Equal(http.StatusOK, rec.Result().StatusCode)
}

func (s *kendangWebhookSuite) TestUnexpectedError() {
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()

	dummyError := errors.New("Unexpected Error")

	mockedSvc := mock.NewMockUpdateSaver(ctrl)
	mockedSvc.EXPECT().Perform(gomock.Any(), gomock.Any()).Return(model.Page{}, dummyError)

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"sourceID":"5d0f6dfbe4e1f617cbbe18b6"}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := s.e.NewContext(req, rec)

	handler := controller.NewKendang(mockedSvc)

	err := handler.OnHandle(c)
	s.Equal(dummyError, err)
}

func TestKendangWebhook(t *testing.T) {
	suite.Run(t, new(kendangWebhookSuite))
}
