package controller

import (
	"github.com/bickyeric/arumba/model"
	"github.com/bickyeric/arumba/service/episode"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Interface ...
type Interface interface {
	OnHandle(c echo.Context) error
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
	data := model.Update{}
	if err := c.Bind(&data); err != nil {
		return err
	}

	sourceID, err := primitive.ObjectIDFromHex(data.SourceID)
	if err != nil {
		return err
	}

	_, err = ctrl.saver.Perform(data, sourceID)
	return err
}
