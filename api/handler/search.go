package handler

import (
	"net/http"

	"github.com/bickyeric/arumba/api"
	"github.com/bickyeric/arumba/service/comic"
	"github.com/labstack/echo"
)

type search struct {
	svc comic.Searcher
}

func NewSearch(svc comic.Searcher) Interface {
	return search{svc}
}

func (h search) OnHandle(c echo.Context) error {
	q := c.QueryParam("q")
	if q == "" {
		return api.NewHTTPError(nil, http.StatusUnprocessableEntity, "q is not provided")
	}

	comics, err := h.svc.Perform(c.Request().Context(), q)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, comics)
}
