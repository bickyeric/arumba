package connection_test

import (
	"testing"

	"github.com/bickyeric/arumba/connection"
	"github.com/stretchr/testify/assert"
)

func TestNewKendangSuccess(t *testing.T) {
	instance := connection.NewKendang()
	assert.NotNil(t, instance)
}
