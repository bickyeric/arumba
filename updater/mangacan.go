package updater

import (
	"github.com/bickyeric/arumba/connection"
	"github.com/bickyeric/arumba/service/episode"
	"github.com/bickyeric/arumba/telegram"
)

type Mangacan struct {
	Bot     telegram.Bot
	Kendang connection.Kendang
	Saver   episode.UpdateSaver
}

func (updater Mangacan) Run() {
	updates, err := updater.Kendang.GetMangacanUpdate()
	if err != nil {
		updater.Bot.NotifyError(err)
		return
	}

	for _, u := range updates {
		err := updater.Saver.Perform(u, 3)
		if err != nil {
			switch err {
			case episode.ErrEpisodeExist:
				continue
			default:
				updater.Bot.NotifyError(err)
			}
		}

		updater.Bot.NotifyNewEpisode(u)
	}
}
