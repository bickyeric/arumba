package model

// Update ...
type Update struct {
	ComicName   string `json:"comicName"`
	EpisodeLink string `json:"episodeLink"`
	EpisodeName string `json:"episodeName"`
	EpisodeNo   int    `json:"episodeNo"`
	SourceID    string `json:"sourceID"`
}
