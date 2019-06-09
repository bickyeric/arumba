package source_test

import (
	"testing"

	"github.com/bickyeric/arumba/updater/source"
	"github.com/stretchr/testify/assert"
)

func TestMangacanName(t *testing.T) {
	s := source.Mangacan{}
	assert.Equal(t, "mangacan", s.Name())
}

func TestMangacanID(t *testing.T) {
	s := source.Mangacan{}
	assert.Equal(t, "5c9511f561a8d04fa844b666", s.GetID().Hex())
}
