package connection

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/bickyeric/arumba/model"
)

type Kendang struct {
	client  *http.Client
	baseURL string
}

func NewKendang() Kendang {
	return Kendang{
		client:  http.DefaultClient,
		baseURL: os.Getenv("KENDANG_URL"),
	}
}

func (k Kendang) GetMangacanUpdate() ([]model.Update, error) {
	result := make([]model.Update, 0)

	request, err := http.NewRequest("GET", k.baseURL+"/mangacan-update", nil)
	if err != nil {
		return nil, err
	}

	response, err := k.client.Do(request)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}
