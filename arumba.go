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

func New(bot telegram.Bot, db *sql.DB) Arumba {
	return Arumba{
		bot:         bot,
		comicRepo:   repository.NewComic(db),
		episodeRepo: repository.NewEpisode(db),
		pageRepo:    repository.NewPage(db),
	}
}

type Arumba struct {
	bot telegram.Bot

	comicRepo   repository.IComic
	episodeRepo repository.IEpisode
	pageRepo    repository.IPage
}

func (kernel Arumba) InjectTelegramStart() command.Start {
	return command.Start{
		Bot: kernel.bot,
		Reader: comic.Read{
			ComicRepo:   kernel.comicRepo,
			EpisodeRepo: kernel.episodeRepo,
			PageRepo:    kernel.pageRepo,
		},
	}
}

func (kernel Arumba) InjectTelegramHelp() command.Help {
	return command.Help{
		Bot: kernel.bot,
	}
}

func (kernel Arumba) InjectTelegramRead() command.Read {
	return command.Read{
		Bot: kernel.bot,
		Reader: comic.Read{
			ComicRepo:   kernel.comicRepo,
			EpisodeRepo: kernel.episodeRepo,
			PageRepo:    kernel.pageRepo,
		},
	}
}

func (kernel Arumba) InjectUpdateRunner() updater.IRunner {
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
