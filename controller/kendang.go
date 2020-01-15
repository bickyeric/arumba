package controller

import (
	"net/http"

	"github.com/bickyeric/arumba/api"
	"github.com/bickyeric/arumba/model"
	"github.com/bickyeric/arumba/service/episode"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Interface ...
type Interface interface {
	OnHandle(echo.Context) error
}

type kendang struct {
	saver episode.UpdateSaver
}

// NewKendang ...
func NewKendang(saver episode.UpdateSaver) Interface {
	return &kendang{
		saver: saver,
	}
}

func (ctrl *kendang) OnHandle(c echo.Context) error {
	var data model.Update
	if err := c.Bind(&data); err != nil {
		return err
	}

	sourceID, err := primitive.ObjectIDFromHex(data.SourceID)
	if err != nil {
		return api.NewHTTPError(err, http.StatusBadRequest, "sourceID is invalid ObjectID")
	}

	page, err := ctrl.saver.Perform(data, sourceID)
	switch err {
	case nil:
		return c.JSON(http.StatusCreated, page)
	case episode.ErrEpisodeExists:
		return c.JSON(http.StatusOK, "ok")
	default:
		return err
	}
}
