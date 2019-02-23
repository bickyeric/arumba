package repository

import (
	"github.com/bickyeric/arumba/connection"
	"github.com/bickyeric/arumba/model"
)

type IEpisode interface {
	FindByNo(comicID int, no float64) (*model.Episode, error)
	GetSources(episodeID int) []int
	Insert(*model.Episode) error
	InsertLink(episodeID, sourceID int, link string) error
	GetLink(episodeID, sourceID int) (string, error)
}

type EpisodeRepository struct{}

func (e EpisodeRepository) InsertLink(episodeID, sourceID int, link string) error {
	_, err := connection.Mysql.Exec("INSERT INTO episode_source(source_id, episode_id, link) VALUES(?,?,?)", sourceID, episodeID, link)
	return err
}

func (e EpisodeRepository) GetLink(episodeID, sourceID int) (string, error) {
	link := ""
	row := connection.Mysql.QueryRow("SELECT link FROM episode_source WHERE source_id=? AND episode_id=?", sourceID, episodeID)
	err := row.Scan(&link)
	return link, err
}

func (r EpisodeRepository) Insert(episode *model.Episode) error {
	res, err := connection.Mysql.Exec("INSERT INTO episodes(no, name, created_at, updated_at, comic_id) VALUES(?,?,?,?,?)", episode.No, episode.Name, episode.CreatedAt, episode.UpdatedAt, episode.ComicID)
	if err != nil {
		return err
	}

	id, _ := res.LastInsertId()
	episode.ID = int(id)
	return nil
}

func (r EpisodeRepository) FindByNo(comicID int, no float64) (*model.Episode, error) {
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
