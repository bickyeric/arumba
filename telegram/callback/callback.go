package callback

// ReadCallback ...
var ReadCallback = "read"

// SelectEpisodeCallback ...
var SelectEpisodeCallback = "select-episode"

// CallbackHandler ...
type CallbackHandler interface {
	Handle(chatID int64, arg string)
}
