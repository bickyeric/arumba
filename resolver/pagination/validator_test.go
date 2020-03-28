package pagination_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/bickyeric/arumba/resolver/pagination"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type validatorSuite struct {
	suite.Suite
}

func (s *validatorSuite) TestBackward_InvalidBeforeCursor() {
	before := "invalid cursor"
	_, err := pagination.Validate(&before, nil, nil, nil)
	s.Equal(pagination.ErrInvalidBeforeCursor, err)
}

func (s *validatorSuite) TestBackward_NegativeLast() {
	before := primitive.NewObjectID().Hex()
	last := -53
	_, err := pagination.Validate(&before, nil, nil, &last)
	s.Equal(pagination.ErrNegativeLast, err)
}

func (s *validatorSuite) TestBackward_LastNotDefined() {
	before := primitive.NewObjectID().Hex()
	p, err := pagination.Validate(&before, nil, nil, nil)
	s.Nil(err)
	jsonPipeline, err := json.Marshal(p.Pipelines())
	s.Nil(err)
	expectedJSON := fmt.Sprintf(`[[{"Key":"$match","Value":{"_id":{"$lt":"%s"}}}],[{"Key":"$sort","Value":{"_id":-1}}],[{"Key":"$limit","Value":5}],[{"Key":"$sort","Value":{"_id":1}}]]`, before)
	s.Equal(expectedJSON, string(jsonPipeline))
}

func (s *validatorSuite) TestBackward_OK() {
	before := primitive.NewObjectID().Hex()
	last := 3
	p, err := pagination.Validate(&before, nil, nil, &last)
	s.Nil(err)
	jsonPipeline, err := json.Marshal(p.Pipelines())
	s.Nil(err)
	expectedJSON := fmt.Sprintf(`[[{"Key":"$match","Value":{"_id":{"$lt":"%s"}}}],[{"Key":"$sort","Value":{"_id":-1}}],[{"Key":"$limit","Value":%d}],[{"Key":"$sort","Value":{"_id":1}}]]`, before, last)
	s.Equal(expectedJSON, string(jsonPipeline))
}

func (s *validatorSuite) TestForward_InvalidAfterCursor() {
	after := "invalid cursor"
	_, err := pagination.Validate(nil, &after, nil, nil)
	s.Equal(pagination.ErrInvalidAfterCursor, err)
}

func (s *validatorSuite) TestForward_NegativeFirst() {
	after := primitive.NewObjectID().Hex()
	first := -8
	_, err := pagination.Validate(nil, &after, &first, nil)
	s.Equal(pagination.ErrNegativeFirst, err)
}

func (s *validatorSuite) TestDefaultPagination() {
	p, err := pagination.Validate(nil, nil, nil, nil)
	s.Nil(err)
	jsonPipeline, err := json.Marshal(p.Pipelines())
	s.Nil(err)
	s.Equal(`[[{"Key":"$limit","Value":5}]]`, string(jsonPipeline))
}

func TestValidator(t *testing.T) {
	suite.Run(t, new(validatorSuite))
}
