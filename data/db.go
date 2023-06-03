package data

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

func InitDB() (db *sql.DB, err error) {
	config := mysql.NewConfig()

	config.Addr = "localhost::3306"
	config.DBName = "todolist"
	config.User = "root"
	config.Passwd = "W3Fm7NpSmhwDqo"

	db, err = sql.Open("mysql", config.FormatDSN())

	return db, err
}
