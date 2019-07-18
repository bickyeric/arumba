package network_test

import (
	"testing"

	"github.com/bickyeric/arumba/connection/network"
	"github.com/stretchr/testify/assert"
)

func TestDecodeValidJSON(t *testing.T) {
	type Model struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Age       int    `json:"age"`
	}

	m := Model{}
	res := network.Response{[]byte(`{
		"first_name" : "bicky eric",
		"last_name" : "kantona",
		"age" : 22
	}`)}

	err := res.Decode(&m)

	assert.Nil(t, err)
	assert.Equal(t, 22, m.Age)
	assert.Equal(t, "bicky eric", m.FirstName)
	assert.Equal(t, "kantona", m.LastName)
}
