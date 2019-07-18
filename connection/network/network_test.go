package network_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bickyeric/arumba/connection/network"
	"github.com/stretchr/testify/assert"
)

func TestNewNetwork(t *testing.T) {
	n := network.New("http://local.host")
	assert.NotNil(t, n)
}

func TestPOST(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)
		res.Write([]byte("FUCK YOU!!!"))
	}))
	defer func() { mockServer.Close() }()

	n := network.New(mockServer.URL + "/")
	res, err := n.POST("", []byte{})
	assert.Nil(t, err)
	assert.Equal(t, "FUCK YOU!!!", string(res.Body))
}

func TestPOST_ErrorDo(t *testing.T) {
	n := network.New("to_invalid_web")
	_, err := n.POST("", []byte{})
	assert.NotNil(t, err)
}

func TestGET(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)
		res.Write([]byte("FUCK YOU!!!"))
	}))
	defer func() { mockServer.Close() }()

	n := network.New(mockServer.URL + "/")
	res, err := n.GET("", []byte{})
	assert.Nil(t, err)
	assert.Equal(t, "FUCK YOU!!!", string(res.Body))
}

func TestGET_ErrorDo(t *testing.T) {
	n := network.New("to_invalid_web")
	_, err := n.GET("", []byte{})
	assert.NotNil(t, err)
}
