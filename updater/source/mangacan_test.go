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
	assert.Equal(t, "5daddd4b73b1d018e959c85b", s.GetID().Hex())
}
