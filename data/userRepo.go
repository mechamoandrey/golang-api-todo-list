package data

import (
	"database/sql"
	"todo-list-api/entities"
)

type userRepo struct {
	db *sql.DB
}

func (r userRepo) CreateUser(uuid string, name string, login string, passwd string) error {
	queryString := `
		INSERT INTO user (user_uuid, name, login, password)
		VALUES (?, ?, ?, ?);
	`

	_, err := r.db.Query(queryString, uuid, name, login, passwd)

	return err
}

func (r userRepo) GetUserByUUID(userUUID string) (user entities.User, err error) {
	queryString := `
		SELECT user_id, user_uuid, name, login, password
		FROM user
		WHERE user_uuid = ?;
	`

	err = r.db.QueryRow(queryString, userUUID).Scan(&user.UserId, &userUUID, &user.Name, &user.Login, &user.Password)

	return user, err
}

func (r userRepo) GetUserByLogin(login string) (user entities.User, err error) {
	queryString := `
		SELECT user_uuid, name, password
		FROM user
		WHERE login = ?;
	`

	err = r.db.QueryRow(queryString, login).Scan(&user.UserUUID, &user.Name, &user.Password)

	return user, err
}
