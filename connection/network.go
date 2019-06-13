package connection

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// ...
var (
	TelegraphNetwork = NewNetwork("https://api.telegra.ph/")
)

// NetworkInterface ...
type NetworkInterface interface {
	POST(link string, body []byte) ([]byte, error)
}

// NewNetwork ...
func NewNetwork(site string) NetworkInterface {
	return network{site, http.DefaultClient}
}

type network struct {
	base   string
	client *http.Client
}

func (n network) POST(link string, bodyParam []byte) ([]byte, error) {
	req, _ := http.NewRequest("POST", n.base+"createPage", bytes.NewBuffer(bodyParam))
	req.Header.Set("Content-Type", "application/json")

	res, err := n.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}
