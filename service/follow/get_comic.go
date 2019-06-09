package follow

import "github.com/bickyeric/arumba/model"

// ComicGetter ...
type ComicGetter interface {
	Perform(chatID int64) ([]model.Comic, error)
}

type getComic struct {
}

// NewComicGetter ...
func NewComicGetter() ComicGetter {
	return getComic{}
}

func (getComic) Perform(chatID int64) ([]model.Comic, error) {
	return nil, nil
}
