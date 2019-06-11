package callback

import (
	"github.com/bickyeric/arumba"
	"github.com/bickyeric/arumba/service/episode"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type selectComic struct {
	bot             arumba.IBot
	notifier        arumba.BotNotifier
	episodeSearcher episode.Searcher
}

// NewSelectComic ...
func NewSelectComic(
	bot arumba.IBot,
	notifier arumba.BotNotifier,
	episodeSearcher episode.Searcher,
) Handler {
	return selectComic{bot, notifier, episodeSearcher}
}

func (handler selectComic) Handle(event *tgbotapi.CallbackQuery) {
	contextLog := log.WithFields(
		log.Fields{
			"chat_id": event.Message.Chat.ID,
		},
	)

	_, arg := ExtractData(event.Data)
	id, _ := primitive.ObjectIDFromHex(arg)
	group, err := handler.episodeSearcher.Perform(id)
	if err != nil {
		handler.notifier.NotifyError(err)
		contextLog.WithFields(
			log.Fields{"error": err},
		).Warn("Error searching episodes")
	}
	handler.bot.SendEpisodeSelector(event.Message.Chat.ID, id, group)
	contextLog.Info("Episode selector sent")
}
