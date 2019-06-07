package callback

// SelectComicCallback ...
var SelectComicCallback = "select-comic"

// SelectEpisodeCallback ...
var SelectEpisodeCallback = "select-episode"

// Handler ...
type Handler interface {
	Handle(chatID int64, arg string)
}
