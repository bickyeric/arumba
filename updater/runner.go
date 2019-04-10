package updater

import (
	"log"

	"github.com/bickyeric/arumba"
	"github.com/bickyeric/arumba/connection"
	"github.com/bickyeric/arumba/service/episode"
	"github.com/bickyeric/arumba/updater/source"
)

// IRunner ...
type IRunner interface {
	Run(source source.ISource)
}

type runner struct {
	bot     arumba.IBot
	kendang connection.IKendang
	saver   episode.UpdateSaver
}

// NewRunner ...
func NewRunner(bot arumba.IBot, kendang connection.IKendang, saver episode.UpdateSaver) IRunner {
	return runner{
		bot:     bot,
		kendang: kendang,
		saver:   saver,
	}
}

// Run ...
func (r runner) Run(source source.ISource) {
	log.Println("Processing " + source.Name() + " updates...")
	updates, err := r.kendang.FetchUpdate("/" + source.Name() + "-update")
	if err != nil {
		r.bot.NotifyError(err)
		return
	}

	for _, u := range updates {
		err := r.saver.Perform(u, source.GetID())
		if err != nil {
			if err.Error() == "episode exists" {
				continue
			} else {
				r.bot.NotifyError(err)
				continue
			}
		}

		r.bot.NotifyNewEpisode(u)
	}
	log.Println(source.Name() + " updates processed.")
}
