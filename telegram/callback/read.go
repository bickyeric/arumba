package callback

import (
	"github.com/bickyeric/arumba"
	"github.com/bickyeric/arumba/service/episode"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ReadHandler ...
type ReadHandler struct {
	Bot             arumba.IBot
	EpisodeSearcher episode.Search
}

func (handler ReadHandler) Handle(chatID int64, arg string) {
	id, _ := primitive.ObjectIDFromHex(arg)
	group, err := handler.EpisodeSearcher.Perform(id)
	if err != nil {
		handler.Bot.NotifyError(err)
	}
	handler.Bot.SendEpisodeSelector(chatID, id, group)
}
