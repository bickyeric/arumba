package comic_test

import (
	"testing"

	"github.com/bickyeric/arumba/resolver/comic"

	"github.com/stretchr/testify/suite"
)

type encodeCursor struct {
	suite.Suite
}

func (s *encodeCursor) TestCursorOK() {
	testcases := []struct {
		cursor string
		value  int
	}{
		{"MQ==", 1},
		{"Mg==", 2},
		{"MjA=", 20},
		{"MjE=", 21},
		{"MjE5", 219},
	}
	for _, tc := range testcases {
		cursor := comic.EncodeCursor(tc.value)
		s.Equal(tc.cursor, cursor)
	}
}

func TestEncodeCursor(t *testing.T) {
	suite.Run(t, new(encodeCursor))
}

type decodeCursor struct {
	suite.Suite
}

func (s *decodeCursor) TestCursorNil() {
	i, err := comic.DecodeCursor(nil)
	s.Nil(err)
	s.Equal(0, i)
}

func (s *decodeCursor) TestCursorIsNotBase64() {
	randomString := "this_is_random_string"
	i, err := comic.DecodeCursor(&randomString)
	s.NotNil(err)
	s.Equal(0, i)
}

func (s *decodeCursor) TestCursorOK() {
	testcases := []struct {
		cursor string
		value  int
	}{
		{"MQ==", 1},
		{"Mg==", 2},
		{"MjA=", 20},
		{"MjE=", 21},
		{"MjE5", 219},
	}
	for _, tc := range testcases {
		i, err := comic.DecodeCursor(&tc.cursor)
		s.Nil(err)
		s.Equal(tc.value, i)
	}
}

func TestDecodeCursor(t *testing.T) {
	suite.Run(t, new(decodeCursor))
}
