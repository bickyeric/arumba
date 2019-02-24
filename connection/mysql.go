package connection

import (
	"database/sql"
	"fmt"
	"os"

	// mysql database driver
	_ "github.com/go-sql-driver/mysql"
)

// NewMysql ...
func NewMysql() *sql.DB {
	var datasource = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", os.Getenv("MYSQL_USERNAME"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"))
	conn, err := sql.Open("mysql", datasource)

	if err != nil {
		panic(err)
	}

	return conn
}
