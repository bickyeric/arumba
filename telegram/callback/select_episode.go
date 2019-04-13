package callback

import (
	"log"

	"github.com/bickyeric/arumba"
	"github.com/bickyeric/arumba/service/episode"
)

// SelectEpisodeHandler ...
type SelectEpisodeHandler struct {
	Bot             arumba.IBot
	EpisodeSearcher episode.Search
}

func (handler SelectEpisodeHandler) Handle(chatID int64, arg string) {
	log.Println(chatID, arg)
	// id, _ := primitive.ObjectIDFromHex(arg)
	// group, err := handler.EpisodeSearcher.Perform(id)
	// if err != nil {
	// 	handler.Bot.NotifyError(err)
	// }
	// handler.Bot.SendEpisodeSelector(chatID, id, group)
}
