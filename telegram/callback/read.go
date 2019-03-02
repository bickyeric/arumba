package callback

import (
	"strconv"

	"github.com/bickyeric/arumba"
	"github.com/bickyeric/arumba/service/episode"
)

// ReadHandler ...
type ReadHandler struct {
	Bot             arumba.IBot
	EpisodeSearcher episode.Search
}

func (handler ReadHandler) Handle(chatID int64, arg string) {
	comicID, err := strconv.Atoi(arg)
	if err != nil {
		handler.Bot.NotifyError(err)
	}

	group, err := handler.EpisodeSearcher.Perform(comicID)
	if err != nil {
		handler.Bot.NotifyError(err)
	}
	handler.Bot.SendEpisodeSelector(chatID, comicID, group)
}
