package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	dbErr  error
)

func init() {
	dbConn, dbErr = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	if dbErr != nil {
		panic(dbErr.Error())
	}
}
