package arumba

import (
	"github.com/bickyeric/arumba/repository"
	"github.com/bickyeric/arumba/service"
	"github.com/bickyeric/arumba/telegram"
	"github.com/bickyeric/arumba/telegram/command"
)

var (
	Instance Arumba
)

func Configure() {
	Instance = Arumba{
		Bot: telegram.BotInstance,

		ComicRepo:   repository.ComicRepository{},
		EpisodeRepo: repository.EpisodeRepository{},
		PageRepo:    repository.PageRepository{},
	}
}

type Arumba struct {
	Bot telegram.Bot

	ComicRepo   repository.ComicRepository
	EpisodeRepo repository.EpisodeRepository
	PageRepo    repository.PageRepository
}

func (kernel Arumba) InjectTelegramStart() command.Start {
	return command.Start{
		Bot: kernel.Bot,
		ComicService: service.ComicService{
			ComicRepo:   kernel.ComicRepo,
			EpisodeRepo: kernel.EpisodeRepo,
			PageRepo:    kernel.PageRepo,
		},
	}
}

func (kernel Arumba) InjectTelegramHelp() command.Help {
	return command.Help{
		Bot: kernel.Bot,
	}
}

func (kernel Arumba) InjectTelegramRead() command.Read {
	return command.Read{
		Bot: kernel.Bot,
		ComicService: service.ComicService{
			ComicRepo:   kernel.ComicRepo,
			EpisodeRepo: kernel.EpisodeRepo,
			PageRepo:    kernel.PageRepo,
		},
	}
}
