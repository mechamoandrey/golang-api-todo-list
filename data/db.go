package data

import (
	"database/sql"
	"todo-list-api/infra/config"

	"github.com/go-sql-driver/mysql"
)

type connection struct {
	db *sql.DB
}

type CnnMask interface {
	ListItemRepo() listItemRepo
	UserRepo() userRepo
}

func (cnn connection) ListItemRepo() listItemRepo {
	return listItemRepo{
		db: cnn.db,
	}
}

func (cnn connection) UserRepo() userRepo {
	return userRepo{
		db: cnn.db,
	}
}

func InitDB(cfg config.Config) (cnn CnnMask, err error) {
	connection := connection{}

	mysqlConfig := mysql.NewConfig()

	mysqlConfig.Addr = cfg.Database.Address + "::" + cfg.Database.Port
	mysqlConfig.DBName = cfg.Database.DBName
	mysqlConfig.User = cfg.Database.User
	mysqlConfig.Passwd = cfg.Database.Passwd

	connection.db, err = sql.Open("mysql", mysqlConfig.FormatDSN())
	if err != nil {
		return connection, err
	}

	err = connection.db.Ping()

	return connection, err
}
