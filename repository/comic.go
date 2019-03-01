package repository

import (
	"database/sql"
	"fmt"

	"github.com/bickyeric/arumba/model"
)

type IComic interface {
	FindOne(name string) (model.Comic, error)
	Search(name string) ([]model.Comic, error)
	Insert(*model.Comic) error
}

type ComicRepository struct {
	*sql.DB
}

func NewComic(db *sql.DB) IComic {
	return ComicRepository{db}
}

func (repo ComicRepository) Insert(comic *model.Comic) error {
	res, err := repo.Exec("INSERT INTO comics(name, status, summary, created_at, updated_at) VALUES(?,?,?,?,?)", comic.Name, "", "", comic.CreatedAt, comic.UpdatedAt)
	if err != nil {
		return err
	}

	id, _ := res.LastInsertId()
	comic.ID = int(id)
	return nil
}

func (repo ComicRepository) FindOne(name string) (model.Comic, error) {
	row := repo.QueryRow(fmt.Sprintf(`SELECT * FROM comics WHERE name LIKE '%%` + name + `%%'`))
	c := model.Comic{}
	summary := sql.NullString{}
	err := row.Scan(&c.ID, &c.Name, &c.Status, &summary, &c.CreatedAt, &c.UpdatedAt)
	if summary.Valid {
		c.Summary = summary.String
	}
	return c, err
}

func (repo ComicRepository) Search(name string) ([]model.Comic, error) {
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
