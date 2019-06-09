package source_test

import (
	"testing"

	"github.com/bickyeric/arumba/updater/source"
	"github.com/stretchr/testify/assert"
)

func TestMangatailName(t *testing.T) {
	s := source.Mangatail{}
	assert.Equal(t, "mangatail", s.Name())
}

func TestMangatailID(t *testing.T) {
	s := source.Mangatail{}
	assert.Equal(t, "5c89e1cb5cff252ae5db8f1e", s.GetID().Hex())
}
