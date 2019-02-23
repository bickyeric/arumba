package model

type Update struct {
	ComicName   string  `json:"comic_name"`
	EpisodeLink string  `json:"episode_link"`
	EpisodeName string  `json:"episode_name"`
	EpisodeNo   float64 `json:"episode_no"`
}
