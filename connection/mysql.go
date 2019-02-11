package connection

import (
	"database/sql"
	"fmt"

	// mysql database driver
	_ "github.com/go-sql-driver/mysql"
)

var Mysql *sql.DB

// Connect ...
func Connect() {
	var datasource = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", "root", "rootpw", "127.0.0.1", 3306, "nbcomic")
	conn, err := sql.Open("mysql", datasource)

	if err != nil {
		panic(err)
	}

	Mysql = conn
}
