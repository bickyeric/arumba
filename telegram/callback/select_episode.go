package callback

import (
	"strconv"
	"strings"

	"github.com/bickyeric/arumba"
	"github.com/bickyeric/arumba/service/comic"
	"github.com/bickyeric/arumba/service/episode"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// SelectEpisodeHandler ...
type SelectEpisodeHandler struct {
	Bot             arumba.IBot
	Notifier        arumba.BotNotifier
	EpisodeSearcher episode.Search
	Reader          comic.Read
}

func (handler SelectEpisodeHandler) Handle(chatID int64, arg string) {
	args := strings.Split(arg, "_")
	if len(args) == 2 {
		handler.readComic(chatID, args)
	} else {
		handler.episodeSelector(chatID, args)
	}
}

func (handler SelectEpisodeHandler) readComic(chatID int64, args []string) {
	contextLog := log.WithFields(
		log.Fields{
			"chat_id": chatID,
		},
	)

	comicID, _ := primitive.ObjectIDFromHex(args[0])
	episodeNo, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		handler.Notifier.NotifyError(err)
		contextLog.WithFields(
			log.Fields{
				"error": err,
			},
		).Warn("Error parsing float")
	}

	pageURL, err := handler.Reader.PerformByComicID(comicID, episodeNo)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			contextLog.WithFields(
				log.Fields{
					"comic_id": comicID,
					"no":       episodeNo,
				},
			).Warn("Episode not found")
			handler.Bot.SendNotFoundEpisode(chatID)
		default:
			contextLog.WithFields(
				log.Fields{
					"comic_id": comicID,
					"no":       episodeNo,
					"error":    err,
				},
			).Warn("Error reading comic")
			handler.Bot.SendErrorMessage(chatID)
		}
		return
	}

	handler.Bot.SendTextMessage(chatID, pageURL)
	contextLog.Info("Page sent")
}

func (handler SelectEpisodeHandler) episodeSelector(chatID int64, args []string) {
	comicID, _ := primitive.ObjectIDFromHex(args[0])
	lower, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		handler.Notifier.NotifyError(err)
	}
	upper, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		handler.Notifier.NotifyError(err)
	}

	group, err := handler.EpisodeSearcher.Perform(comicID, lower, upper)
	if err != nil {
		handler.Notifier.NotifyError(err)
	}
	handler.Bot.SendEpisodeSelector(chatID, comicID, group)
}
