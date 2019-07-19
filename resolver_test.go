package arumba_test

import (
	"testing"

	"github.com/bickyeric/arumba"
	"github.com/stretchr/testify/assert"
)

func TestResolver(t *testing.T) {
	r := arumba.Resolver{nil, nil, nil}

	assert.NotPanics(t, func() {
		r.Comic()
	})

	assert.NotPanics(t, func() {
		r.Episode()
	})

	assert.NotPanics(t, func() {
		r.Source()
	})

	assert.NotPanics(t, func() {
		r.Query()
	})
}
