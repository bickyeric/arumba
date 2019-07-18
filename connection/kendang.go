package connection

import (
	"fmt"
	"os"

	"github.com/bickyeric/arumba/connection/network"
	"github.com/bickyeric/arumba/model"
	log "github.com/sirupsen/logrus"
)

// IKendang ...
type IKendang interface {
	FetchUpdate(source string) ([]model.Update, error)
	FetchPages(episodeLink string, sourceID string) ([]string, error)
}

type kendang struct {
	network network.Interface
}

// NewKendang ...
func NewKendang() IKendang {
	return kendang{
		network.New(os.Getenv("KENDANG_URL")),
	}
}

func (k kendang) FetchUpdate(source string) ([]model.Update, error) {
	result := make([]model.Update, 0)

	response, err := k.network.GET(source, nil)
	if err != nil {
		return nil, err
	}

	if err := response.Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func (k kendang) FetchPages(episodeLink string, sourceID string) ([]string, error) {
	id := k.toID(sourceID)
	link := fmt.Sprintf("crawl-page?link=%s&source_id=%d", episodeLink, id)

	log.WithFields(
		log.Fields{
			"link": link,
		},
	).Info("Crawling page from kendang")
	response, err := k.network.GET(link, nil)
	if err != nil {
		return nil, err
	}

	pagesLink := []string{}
	err = response.Decode(&pagesLink)
	return pagesLink, err
}

func (kendang) toID(sourceID string) int {
	switch sourceID {
	case "5c9511f561a8d04fa844b666":
		return 3
	case "5c89e1cb5cff252ae5db8f1e":
		return 2
	}
	return 0
}
