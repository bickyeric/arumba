package controller_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bickyeric/arumba/controller"
	"github.com/bickyeric/arumba/service/episode/mock"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestKendang(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	t.Run("invalid json body", func(t *testing.T) {
		mockedSvc := mock.NewMockUpdateSaver(ctrl)

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{"))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		assert.NotPanics(t, func() {
			handler := controller.NewKendang(mockedSvc)

			err := handler.OnHandle(c)
			assert.NotNil(t, err)
		})
	})

	t.Run("invalid sourceID", func(t *testing.T) {
		mockedSvc := mock.NewMockUpdateSaver(ctrl)

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"sourceID":"hehe"}`))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		assert.NotPanics(t, func() {
			handler := controller.NewKendang(mockedSvc)

			err := handler.OnHandle(c)
			assert.NotNil(t, err)
		})
	})

	t.Run("ok", func(t *testing.T) {
		mockedSvc := mock.NewMockUpdateSaver(ctrl)
		mockedSvc.EXPECT().Perform(gomock.Any(), gomock.Any())

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"sourceID":"5d0f6dfbe4e1f617cbbe18b6"}`))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		assert.NotPanics(t, func() {
			handler := controller.NewKendang(mockedSvc)

			err := handler.OnHandle(c)
			assert.Nil(t, err)
		})
	})
}
