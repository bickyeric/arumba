package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/bickyeric/arumba/connection"
)

// Comic merepresentasikan objek komik
type Comic struct {
	Source  Source  `json:"source"`
	Episode Episode `json:"episode"`
}

func ReadComic(name string, episode int) (*Comic, error) {
	dataSource := connection.DataSourceInstance()
	url := fmt.Sprintf("%s/api/comic/%s/episode/%d", dataSource.BaseURL, strings.ToLower(name), episode)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := dataSource.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	jsonRaw, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	dec := json.NewDecoder(strings.NewReader(string(jsonRaw)))
	var comic Comic
	if err := dec.Decode(&comic); err != nil {
		return nil, err
	}

	return &comic, nil
}
