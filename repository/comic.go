package repository

import (
	"database/sql"
	"fmt"

	"github.com/bickyeric/arumba/connection"
	"github.com/bickyeric/arumba/model"
)

type IComic interface {
	FindByName(name string) (*model.Comic, error)
	Insert(*model.Comic) error
}

type ComicRepository struct{}

func (r ComicRepository) Insert(comic *model.Comic) error {
	res, err := connection.Mysql.Exec("INSERT INTO comics(name, status, summary, created_at, updated_at) VALUES(?,?,?,?,?)", comic.Name, "", "", comic.CreatedAt, comic.UpdatedAt)
	if err != nil {
		return err
	}

	id, _ := res.LastInsertId()
	comic.ID = int(id)
	return nil
}

func (r ComicRepository) FindByName(name string) (*model.Comic, error) {
	comic := new(model.Comic)
	query := fmt.Sprintf(`SELECT * FROM comics WHERE name LIKE '%%` + name + `%%'`)
	row := connection.Mysql.QueryRow(query)
	summary := sql.NullString{}
	err := row.Scan(&comic.ID, &comic.Name, &comic.Status, &summary, &comic.CreatedAt, &comic.UpdatedAt)
	if summary.Valid {
		comic.Summary = summary.String
	}
	return comic, err
}
