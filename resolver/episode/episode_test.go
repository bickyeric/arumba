package episode_test

import (
	"testing"

	"github.com/bickyeric/arumba/resolver/episode"
	"github.com/stretchr/testify/assert"
)

func TestEncodeCursor(t *testing.T) {
	testCase := []struct {
		no      int
		encoded string
	}{
		{1, "MQ=="},
		{2, "Mg=="},
		{3, "Mw=="},
		{4, "NA=="},
		{5, "NQ=="},
		{10, "MTA="},
		{891, "ODkx"},
	}
	for _, tt := range testCase {
		assert.Equal(t, tt.encoded, episode.EncodeCursor(tt.no))
	}
}

func TestDecodeCursor(t *testing.T) {
	testCase := []struct {
		no      int
		encoded string
	}{
		{1, "MQ=="},
		{2, "Mg=="},
		{3, "Mw=="},
		{4, "NA=="},
		{5, "NQ=="},
		{10, "MTA="},
		{891, "ODkx"},
	}
	for _, tt := range testCase {
		no, err := episode.DecodeCursor(tt.encoded)
		assert.Nil(t, err)
		assert.Equal(t, tt.no, no)
	}
}

func TestDecodeCursor_InvalidCursor(t *testing.T) {
	invalidCursor := "this is invalid cursor"
	_, err := episode.DecodeCursor(invalidCursor)
	assert.NotNil(t, err)
}
