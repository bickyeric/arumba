package model_test

import (
	"bytes"
	"testing"

	"github.com/bickyeric/arumba/model"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type modelSuite struct {
	suite.Suite
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
