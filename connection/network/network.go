package network

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// Interface ...
type Interface interface {
	POST(link string, body []byte) (Response, error)
	GET(link string, body []byte) (Response, error)
}

// New ...
func New(site string) Interface {
	return network{site, http.DefaultClient}
}

type network struct {
	base   string
	client *http.Client
}

func (n network) POST(link string, bodyParam []byte) (Response, error) {
	res := Response{}
	req, _ := http.NewRequest("POST", n.base+link, bytes.NewBuffer(bodyParam))
	req.Header.Set("Content-Type", "application/json")

	httpres, err := n.client.Do(req)
	if err != nil {
		return res, err
	}
	defer httpres.Body.Close()

	res.Body, err = ioutil.ReadAll(httpres.Body)
	return res, err
}

func (n network) GET(link string, bodyParam []byte) (Response, error) {
	res := Response{}
	req, _ := http.NewRequest("GET", n.base+link, bytes.NewBuffer(bodyParam))
	req.Header.Set("Content-Type", "application/json")

	httpres, err := n.client.Do(req)
	if err != nil {
		return res, err
	}
	defer httpres.Body.Close()

	res.Body, err = ioutil.ReadAll(httpres.Body)
	return res, err
}
