package callback

import (
	"database/sql"
	"strconv"
	"strings"

	"github.com/bickyeric/arumba"
	"github.com/bickyeric/arumba/service/comic"
	"github.com/bickyeric/arumba/service/episode"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SelectEpisodeHandler ...
type SelectEpisodeHandler struct {
	Bot             arumba.IBot
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
	comicID, _ := primitive.ObjectIDFromHex(args[0])
	episodeNo, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		handler.Bot.NotifyError(err)
	}
	pages, err := handler.Reader.PerformByComicID(comicID, episodeNo)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			handler.Bot.SendNotFoundEpisode(chatID)
		default:
			handler.Bot.SendErrorMessage(chatID)
		}
		return
	}

	handler.Bot.SendPage(chatID, pages)
}

func (handler SelectEpisodeHandler) episodeSelector(chatID int64, args []string) {
	comicID, _ := primitive.ObjectIDFromHex(args[0])
	lower, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		handler.Bot.NotifyError(err)
	}
	upper, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		handler.Bot.NotifyError(err)
	}

	group, err := handler.EpisodeSearcher.Perform(comicID, lower, upper)
	if err != nil {
		handler.Bot.NotifyError(err)
	}
	handler.Bot.SendEpisodeSelector(chatID, comicID, group)
}
