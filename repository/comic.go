package repository

import (
	"database/sql"
	"fmt"

	"github.com/bickyeric/arumba/model"
)

// IComic ...
type IComic interface {
	Find(name string) (model.Comic, error)
	FindAll(name string) ([]model.Comic, error)
	Insert(*model.Comic) error
}

type comicRepository struct {
	*sql.DB
}

// NewComic ...
func NewComic(db *sql.DB) IComic {
	return comicRepository{db}
}

func (repo comicRepository) Insert(comic *model.Comic) error {
	res, err := repo.Exec("INSERT INTO comics(name, status, summary, created_at, updated_at) VALUES(?,?,?,?,?)", comic.Name, "", "", comic.CreatedAt, comic.UpdatedAt)
	if err != nil {
		return err
	}

	id, _ := res.LastInsertId()
	comic.ID = int(id)
	return nil
}

func (repo comicRepository) Find(name string) (model.Comic, error) {
	row := repo.QueryRow(fmt.Sprintf(`SELECT * FROM comics WHERE name LIKE '%%` + name + `%%'`))
	c := model.Comic{}
	summary := sql.NullString{}
	err := row.Scan(&c.ID, &c.Name, &c.Status, &summary, &c.CreatedAt, &c.UpdatedAt)
	if summary.Valid {
		c.Summary = summary.String
	}
	return c, err
}

func (repo comicRepository) FindAll(name string) ([]model.Comic, error) {
	row, err := repo.Query(fmt.Sprintf(`SELECT * FROM comics WHERE name LIKE '%%` + name + `%%'`))
	if err != nil {
		return nil, err
	}

	comics := []model.Comic{}
	for row.Next() {
		c := model.Comic{}
		summary := sql.NullString{}
		row.Scan(&c.ID, &c.Name, &c.Status, &summary, &c.CreatedAt, &c.UpdatedAt)
		if summary.Valid {
			c.Summary = summary.String
		}
		comics = append(comics, c)
	}
	return comics, err
}
