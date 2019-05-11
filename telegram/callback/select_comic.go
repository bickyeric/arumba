package callback

import (
	"github.com/bickyeric/arumba"
	"github.com/bickyeric/arumba/service/episode"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SelectComicHandler ...
type SelectComicHandler struct {
	Bot             arumba.IBot
	EpisodeSearcher episode.Search
}

func (handler SelectComicHandler) Handle(chatID int64, arg string) {
	contextLog := log.WithFields(
		log.Fields{
			"chat_id": chatID,
		},
	)

	id, _ := primitive.ObjectIDFromHex(arg)
	group, err := handler.EpisodeSearcher.Perform(id)
	if err != nil {
		handler.Bot.NotifyError(err)
		contextLog.WithFields(
			log.Fields{"error": err},
		).Warn("Error searching episodes")
	}
	handler.Bot.SendEpisodeSelector(chatID, id, group)
	contextLog.Info("Episode selector sent")
}
