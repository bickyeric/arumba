package repository

import (
	"github.com/bickyeric/arumba/connection"
	"github.com/bickyeric/arumba/model"
)

type IEpisode interface {
	FindByNo(comicID, no int) (*model.Episode, error)
	GetSources(episodeID int) []int
}

type EpisodeRepository struct{}

func (r EpisodeRepository) FindByNo(comicID, no int) (*model.Episode, error) {
	episode := new(model.Episode)
	row := connection.Mysql.QueryRow("SELECT * FROM episodes WHERE comic_id=? AND no=?", comicID, no)
	err := row.Scan(&episode.ID, &episode.No, &episode.Name, &episode.CreatedAt, &episode.UpdatedAt, &episode.ComicID)
	return episode, err
}

func (r EpisodeRepository) GetSources(episodeID int) []int {
	sourceIds := []int{}
	rows, err := connection.Mysql.Query("SELECT source_id FROM episode_source WHERE episode_id=?", episodeID)
	if err != nil {
		return sourceIds
	}
	for rows.Next() {
		var id int
		rows.Scan(&id)
		sourceIds = append(sourceIds, id)
	}
	return sourceIds
}
