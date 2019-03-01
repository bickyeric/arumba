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
	InjectTelegramStart() command.Start
	InjectTelegramHelp() command.Help
	InjectTelegramRead() command.Read
	InjectTelegramFeedback() command.Feedback
	InjectTelegramCommon() command.Common

	InjectUpdateRunner() updater.IRunner
}

// New ...
func New(bot telegram.Bot, db *sql.DB) IArumba {
	return arumba{
		bot:         bot,
		comicRepo:   repository.NewComic(db),
		episodeRepo: repository.NewEpisode(db),
		pageRepo:    repository.NewPage(db),
	}
}

type arumba struct {
	bot telegram.Bot

	comicRepo   repository.IComic
	episodeRepo repository.IEpisode
	pageRepo    repository.IPage
}

func (kernel arumba) InjectTelegramStart() command.Start {
	return command.Start{
		Bot: kernel.bot,
		Reader: comic.Read{
			ComicRepo:   kernel.comicRepo,
			EpisodeRepo: kernel.episodeRepo,
			PageRepo:    kernel.pageRepo,
			Kendang:     connection.NewKendang(),
		},
	}
}

func (kernel arumba) InjectTelegramHelp() command.Help {
	return command.Help{
		Bot: kernel.bot,
	}
}

func (kernel arumba) InjectTelegramRead() command.Read {
	return command.Read{
		Bot: kernel.bot,
		Reader: comic.Read{
			ComicRepo:   kernel.comicRepo,
			EpisodeRepo: kernel.episodeRepo,
			PageRepo:    kernel.pageRepo,
		},
	}
}

func (kernel arumba) InjectTelegramFeedback() command.Feedback {
	return command.Feedback{
		Bot: kernel.bot,
	}
}

func (kernel arumba) InjectTelegramCommon() command.Common {
	return command.Common{
		Bot: kernel.bot,
		ComicSearcher: comic.Search{
			Repo: kernel.comicRepo,
		},
	}
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
