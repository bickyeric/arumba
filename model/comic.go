package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/bickyeric/arumba/connection"
)

// Comic merepresentasikan objek komik
type Comic struct {
	Source  Source  `json:"source"`
	Episode Episode `json:"episode"`
}

var readPath = "/comic/read/%d/%s/%s"

func ReadComic(name string, episode int, userID int64) *Comic {
	dataSource := connection.DataSourceInstance()
	url := fmt.Sprintf(dataSource.BaseURL+readPath, userID, name, episode)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := dataSource.HTTPClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	jsonRaw, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	dec := json.NewDecoder(strings.NewReader(string(jsonRaw)))
	var comic Comic
	if err := dec.Decode(&comic); err != nil {
		log.Fatal(err)
	}

	return &comic
}
