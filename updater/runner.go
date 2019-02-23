package updater

import (
	"log"

	"github.com/bickyeric/arumba/connection"
	"github.com/bickyeric/arumba/service/episode"
	"github.com/bickyeric/arumba/telegram"
)

// Runner ...
type Runner struct {
	Bot     telegram.Bot
	Kendang connection.Kendang
	Saver   episode.UpdateSaver
}

// Run ...
func (r Runner) Run(source ISource) {
	log.Println("Processing " + source.Name() + " updates...")
	updates, err := r.Kendang.FetchUpdate("/" + source.Name() + "-update")
	if err != nil {
		r.Bot.NotifyError(err)
		return
	}

	for _, u := range updates {
		err := r.Saver.Perform(u, source.GetID())
		if err != nil {
			switch err {
			case episode.ErrEpisodeExist:
				continue
			default:
				r.Bot.NotifyError(err)
			}
		}

		r.Bot.NotifyNewEpisode(u)
	}
	log.Println(source.Name() + " updates processed.")
}
