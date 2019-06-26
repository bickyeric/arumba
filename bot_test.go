package arumba_test

import (
	"testing"

	"github.com/bickyeric/arumba"
)

func TestNotifyNewEpisode(t *testing.T) {
	bot := arumba.Bot{nil}
	bot.NotifyNewEpisode("One Piece", "http://localhost/one-piece/890", 890)
}
