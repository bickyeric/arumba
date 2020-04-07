package pagination_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/bickyeric/arumba/resolver/pagination"
	"github.com/stretchr/testify/suite"
)

type validatorSuite struct {
	suite.Suite
}

func (s *validatorSuite) TestForward_InvalidAfterCursor() {
	after := "invalid cursor"
	_, err := pagination.Validate(&after, nil)
	s.Equal(pagination.ErrInvalidAfterCursor, err)
}

func (s *validatorSuite) TestForward_NegativeFirst() {
	after := "ODkx"
	first := -8
	_, err := pagination.Validate(&after, &first)
	s.Equal(pagination.ErrNegativeFirst, err)
}

func (s *validatorSuite) TestForward_FirstNotDefined() {
	after := "ODkx"
	p, err := pagination.Validate(&after, nil)
	s.Nil(err)
	jsonPipeline, err := json.Marshal(p.Pipelines())
	s.Nil(err)
	expectedJSON := fmt.Sprintf(`[[{"Key":"$match","Value":{"no":{"$lt":%d}}}],[{"Key":"$limit","Value":%d}]]`, 891, 5)
	s.Equal(expectedJSON, string(jsonPipeline))
}

func (s *validatorSuite) TestForward_OK() {
	after := "ODkx"
	first := 10
	p, err := pagination.Validate(&after, &first)
	s.Nil(err)
	jsonPipeline, err := json.Marshal(p.Pipelines())
	s.Nil(err)
	expectedJSON := fmt.Sprintf(`[[{"Key":"$match","Value":{"no":{"$lt":%d}}}],[{"Key":"$limit","Value":%d}]]`, 891, first)
	s.Equal(expectedJSON, string(jsonPipeline))
}

func (s *validatorSuite) TestDefaultPagination() {
	p, err := pagination.Validate(nil, nil)
	s.Nil(err)
	jsonPipeline, err := json.Marshal(p.Pipelines())
	s.Nil(err)
	s.Equal(`[[{"Key":"$limit","Value":5}]]`, string(jsonPipeline))
}

func TestValidator(t *testing.T) {
	suite.Run(t, new(validatorSuite))
}
