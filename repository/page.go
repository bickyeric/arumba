package repository

import (
	"database/sql"

	"github.com/bickyeric/arumba/model"
)

type IPage interface {
	FindByEpisode(episodeID, sourceID int) ([]*model.Page, error)
	Insert(*model.Page) error
}

type PageRepository struct {
	*sql.DB
}

func NewPage(db *sql.DB) IPage {
	return PageRepository{db}
}

func (repo PageRepository) FindByEpisode(episodeID, sourceID int) ([]*model.Page, error) {
	result := []*model.Page{}

	rows, err := repo.Query(`SELECT * FROM pages WHERE episode_id=? AND source_id=?`, episodeID, sourceID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		page := new(model.Page)
		err = rows.Scan(&page.ID, &page.Link, &page.EpisodeID, &page.SourceID)
		if err != nil {
			return nil, err
		}
		result = append(result, page)
	}
	return result, nil
}

func (repo PageRepository) Insert(page *model.Page) error {
	result, err := repo.Exec("INSERT INTO pages(link, episode_id, source_id) VALUES(?,?,?)", page.Link, page.EpisodeID, page.SourceID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	page.ID = int(id)
	return nil
}
