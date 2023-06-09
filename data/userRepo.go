package data

import (
	"database/sql"
)

type userRepo struct {
	db *sql.DB
}

func (r userRepo) HandleRegisterUser(uuid string, name string, login string, passwd string) error {
	queryString := `
		INSERT INTO user (user_uuid, name, login, password)
		VALUES (?, ?, ?, ?);
	`

	_, err := r.db.Query(queryString, uuid, name, login, passwd)

	return err
}
