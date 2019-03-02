package arumba

import (
	"database/sql"

	"github.com/bickyeric/arumba/connection"
	"github.com/bickyeric/arumba/repository"
	"github.com/bickyeric/arumba/service/comic"
	"github.com/bickyeric/arumba/service/episode"
	"github.com/bickyeric/arumba/telegram"
	"github.com/bickyeric/arumba/telegram/command"
	"github.com/bickyeric/arumba/updater"
)

// IArumba ...
type IArumba interface {
	InjectTelegramStart() telegram.CommandHandler
	InjectTelegramHelp() telegram.CommandHandler
	InjectTelegramRead() telegram.CommandHandler
	InjectTelegramFeedback() telegram.CommandHandler
	InjectTelegramCommon() telegram.CommandHandler

	InjectUpdateRunner() updater.IRunner
}

// New ...
func New(bot telegram.IBot, db *sql.DB) IArumba {
	return arumba{
		bot:         bot,
		comicRepo:   repository.NewComic(db),
		episodeRepo: repository.NewEpisode(db),
		pageRepo:    repository.NewPage(db),
	}
}

type arumba struct {
	bot telegram.IBot

	comicRepo   repository.IComic
	episodeRepo repository.IEpisode
	pageRepo    repository.IPage
}

func (kernel arumba) InjectTelegramStart() telegram.CommandHandler {
	return command.Start(
		kernel.bot,
		comic.Read{
			ComicRepo:   kernel.comicRepo,
			EpisodeRepo: kernel.episodeRepo,
			PageRepo:    kernel.pageRepo,
			Kendang:     connection.NewKendang(),
		},
	)
}

func (kernel arumba) InjectTelegramHelp() telegram.CommandHandler {
	return command.Help(kernel.bot)
}

func (kernel arumba) InjectTelegramRead() telegram.CommandHandler {
	return command.Read(
		kernel.bot,
		comic.Read{
			ComicRepo:   kernel.comicRepo,
			EpisodeRepo: kernel.episodeRepo,
			PageRepo:    kernel.pageRepo,
		},
	)
}

func (kernel arumba) InjectTelegramFeedback() telegram.CommandHandler {
	return command.Feedback(kernel.bot)
}

func (kernel arumba) InjectTelegramCommon() telegram.CommandHandler {
	return command.Common(
		kernel.bot,
		comic.Search{
			Repo: kernel.comicRepo,
		},
	)
}

func (kernel arumba) InjectUpdateRunner() updater.IRunner {
	return updater.NewRunner(
		kernel.bot,
		connection.NewKendang(),
		episode.UpdateSaver{
			ComicRepo:   kernel.comicRepo,
			EpisodeRepo: kernel.episodeRepo,
			PageRepo:    kernel.pageRepo,
		},
	)
}
