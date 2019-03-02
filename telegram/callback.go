package telegram

// ReadCallback ...
var ReadCallback = "read"

// CallbackHandler ...
type CallbackHandler interface {
	Handle(chatID int64, arg string)
}
