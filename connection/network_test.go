package connection_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bickyeric/arumba/connection"
	"github.com/stretchr/testify/assert"
)

func TestNewNetwork(t *testing.T) {
	n := connection.NewNetwork("http://local.host")
	assert.NotNil(t, n)
}

func TestPOST(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)
		res.Write([]byte("FUCK YOU!!!"))
	}))
	defer func() { mockServer.Close() }()

	n := connection.NewNetwork(mockServer.URL + "/")
	res, err := n.POST("", []byte{})
	assert.Nil(t, err)
	assert.Equal(t, "FUCK YOU!!!", string(res))
}

func TestPOST_ErrorDo(t *testing.T) {
	n := connection.NewNetwork("to_invalid_web")
	_, err := n.POST("", []byte{})
	assert.NotNil(t, err)
}
