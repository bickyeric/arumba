package model_test

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/bickyeric/arumba/model"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type modelSuite struct {
	suite.Suite
}

func (s *modelSuite) TestUnmarshalTimestamp() {
	expected := time.Now()
	actual, err := model.UnmarshalTimestamp(expected.Unix())
	s.Nil(err)
	s.Equal(expected.Year(), actual.Year())
	s.Equal(expected.Month(), actual.Month())
	s.Equal(expected.Day(), actual.Day())
	s.Equal(expected.Hour(), actual.Hour())
	s.Equal(expected.Minute(), actual.Minute())
	s.Equal(expected.Second(), actual.Second())

	_, err = model.UnmarshalTimestamp("invalid id")
	s.NotNil(err)
}

func (s *modelSuite) TestMarshalTimestamp() {
	var b bytes.Buffer
	expected := time.Now()

	m := model.MarshalTimestamp(expected)
	m.MarshalGQL(&b)
	s.Equal(fmt.Sprintf("%d", expected.Unix()), b.String())
}

func (s *modelSuite) TestUnmarshalID() {
	expected := primitive.NewObjectID()
	actual, err := model.UnmarshalID(expected.Hex())
	s.Nil(err)
	s.Equal(expected, actual)

	_, err = model.UnmarshalID("invalid id")
	s.NotNil(err)

	_, err = model.UnmarshalID(1)
	s.NotNil(err)
}

func (s *modelSuite) TestMarshalID() {
	var b bytes.Buffer
	expected := primitive.NewObjectID()
	m := model.MarshalID(expected)
	m.MarshalGQL(&b)
	s.Equal("\""+expected.Hex()+"\"", b.String())
}

func TestRootResolver(t *testing.T) {
	suite.Run(t, new(modelSuite))
}
