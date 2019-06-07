package connection_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/bickyeric/arumba/connection"
	"github.com/stretchr/testify/assert"
)

func TestNewKendangSuccess(t *testing.T) {
	instance := connection.NewKendang()
	assert.NotNil(t, instance)
}

func TestFetchUpdate_InvalidRequest(t *testing.T) {

	os.Setenv("KENDANG_URL", "this_is_invalid_url")

	instance := connection.NewKendang()
	_, err := instance.FetchUpdate("mangacan")
	assert.NotNil(t, err)
}

func TestFetchUpdate_InvalidResponse(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)
		res.Write([]byte("this_is_invalid_response_from_server"))
	}))
	defer func() { mockServer.Close() }()

	os.Setenv("KENDANG_URL", mockServer.URL+"/")

	instance := connection.NewKendang()
	_, err := instance.FetchUpdate("mangacan")
	assert.NotNil(t, err)
}

func TestFetchUpdate(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)
		res.Write([]byte("[]"))
	}))
	defer func() { mockServer.Close() }()

	os.Setenv("KENDANG_URL", mockServer.URL+"/")

	instance := connection.NewKendang()
	updates, err := instance.FetchUpdate("mangacan")
	assert.Nil(t, err)
	assert.Len(t, updates, 0)
}

func TestFetchPages_InvalidRequest(t *testing.T) {
	os.Setenv("KENDANG_URL", "this_is_invalid_url")

	instance := connection.NewKendang()
	_, err := instance.FetchPages("http://www.mangacanblog.com/baca-komik-one_piece-942-943-bahasa-indonesia-one_piece-942-terbaru.html", "5c9511f561a8d04fa844b666")
	assert.NotNil(t, err)
}

func TestFetchPages_InvalidResponse(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)
		res.Write([]byte("this_is_invalid_response_from_server"))
	}))
	defer func() { mockServer.Close() }()

	os.Setenv("KENDANG_URL", mockServer.URL+"/")

	instance := connection.NewKendang()
	_, err := instance.FetchPages("http://www.mangacanblog.com/baca-komik-one_piece-942-943-bahasa-indonesia-one_piece-942-terbaru.html", "5c89e1cb5cff252ae5db8f1e")
	assert.NotNil(t, err)
}

func TestFetchPages(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)
		res.Write([]byte("[]"))
	}))
	defer func() { mockServer.Close() }()

	os.Setenv("KENDANG_URL", mockServer.URL+"/")

	instance := connection.NewKendang()
	_, err := instance.FetchPages("http://www.mangacanblog.com/baca-komik-one_piece-942-943-bahasa-indonesia-one_piece-942-terbaru.html", "5ds")
	assert.Nil(t, err)
}
