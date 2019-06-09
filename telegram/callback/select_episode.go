package callback

import (
	"strconv"
	"strings"

	"github.com/bickyeric/arumba"
	"github.com/bickyeric/arumba/service/comic"
	"github.com/bickyeric/arumba/service/episode"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type selectEpisode struct {
	bot             arumba.IBot
	notifier        arumba.BotNotifier
	episodeSearcher episode.Search
	reader          comic.Read
}

// NewSelectEpisode ...
func NewSelectEpisode(
	bot arumba.IBot,
	notifier arumba.BotNotifier,
	episodeSearcher episode.Search,
	reader comic.Read,
) Handler {
	return selectEpisode{bot, notifier, episodeSearcher, reader}
}

func (handler selectEpisode) Handle(event *tgbotapi.CallbackQuery) {
	_, arg := ExtractData(event.Data)
	args := strings.Split(arg, "_")
	if len(args) == 2 {
		handler.readComic(event.Message.Chat.ID, args)
	} else {
		handler.episodeSelector(event.Message.Chat.ID, args)

	}
}

func (handler selectEpisode) readComic(chatID int64, args []string) {
	contextLog := log.WithFields(
		log.Fields{
			"chat_id": chatID,
		},
	)

	comicID, _ := primitive.ObjectIDFromHex(args[0])
	episodeNo, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		handler.notifier.NotifyError(err)
		contextLog.WithFields(
			log.Fields{
				"error": err,
			},
		).Warn("Error parsing float")
	}

	pageURL, err := handler.reader.PerformByComicID(comicID, episodeNo)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			contextLog.WithFields(
				log.Fields{
					"comic_id": comicID,
					"no":       episodeNo,
				},
			).Warn("Episode not found")
			handler.bot.SendNotFoundEpisode(chatID)
		default:
			contextLog.WithFields(
				log.Fields{
					"comic_id": comicID,
					"no":       episodeNo,
					"error":    err,
				},
			).Warn("Error reading comic")
			handler.bot.SendErrorMessage(chatID)
		}
		return
	}

	handler.bot.SendTextMessage(chatID, pageURL)
	contextLog.Info("Page sent")
}

func (handler selectEpisode) episodeSelector(chatID int64, args []string) {
	comicID, _ := primitive.ObjectIDFromHex(args[0])
	lower, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		handler.notifier.NotifyError(err)
	}
	upper, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		handler.notifier.NotifyError(err)
	}

	group, err := handler.episodeSearcher.Perform(comicID, lower, upper)
	if err != nil {
		handler.notifier.NotifyError(err)
	}
	handler.bot.SendEpisodeSelector(chatID, comicID, group)
}
