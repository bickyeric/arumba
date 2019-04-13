package callback

// SelectComic ...
var SelectComicCallback = "select-comic"

// SelectEpisodeCallback ...
var SelectEpisodeCallback = "select-episode"

// CallbackHandler ...
type CallbackHandler interface {
	Handle(chatID int64, arg string)
}
