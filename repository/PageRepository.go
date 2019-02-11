package repository

import (
	"github.com/bickyeric/arumba/connection"
	"github.com/bickyeric/arumba/model"
)

type IPage interface {
	FindByEpisode(episodeID, sourceID int) ([]*model.Page, error)
}

type PageRepository struct{}

func (r PageRepository) FindByEpisode(episodeID, sourceID int) ([]*model.Page, error) {
	result := []*model.Page{}

	rows, err := connection.Mysql.Query(`SELECT * FROM pages WHERE episode_id=? AND source_id=?`, episodeID, sourceID)
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
