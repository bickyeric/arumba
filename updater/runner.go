package updater

import (
	"time"

	"github.com/bickyeric/arumba"
	"github.com/bickyeric/arumba/connection"
	"github.com/bickyeric/arumba/service/episode"
	"github.com/bickyeric/arumba/updater/source"
	log "github.com/sirupsen/logrus"
)

// IRunner ...
type IRunner interface {
	Run(source source.ISource)
}

type runner struct {
	notifier arumba.BotNotifier
	kendang  connection.IKendang
	saver    episode.UpdateSaver
}

// NewRunner ...
func NewRunner(bot arumba.BotNotifier, kendang connection.IKendang, app arumba.Arumba) IRunner {
	return runner{
		notifier: bot,
		kendang:  kendang,
		saver:    episode.NewSaveUpdate(app, kendang),
	}
}

// Run ...
func (r runner) Run(source source.ISource) {
	start := time.Now()
	contextLogger := log.WithFields(log.Fields{
		"source": source.Name(),
	})
	contextLogger.Info("Processing updates")

	updates, err := r.kendang.FetchUpdate("/" + source.Name() + "-update")
	if err != nil {
		r.notifier.NotifyError(err)
		contextLogger.WithFields(log.Fields{
			"error": err.Error(),
		}).Warn("Error fetching updates")
		return
	}

	for _, u := range updates {
		_, err := r.saver.Perform(u, source.GetID())
		if err != nil {
			if err == episode.ErrEpisodeExists {
				continue
			} else {
				r.notifier.NotifyError(err)
				contextLogger.WithFields(log.Fields{
					"update": u,
					"error":  err.Error(),
				}).Warn("Error processing update")
				continue
			}
		}

		r.notifier.NotifyNewEpisode(u.ComicName, u.EpisodeLink, int(u.EpisodeNo))
	}

	elapsed := time.Since(start)
	contextLogger.WithFields(log.Fields{
		"duration": elapsed.Seconds(),
	}).Info("Updates processed")
}
