package data

import (
	"database/sql"
	"log"
	"todo-list-api/entities"
)

type listItemRepo struct {
	db *sql.DB
}

func (r listItemRepo) GetListItemByUserId(userUUID string) (userList []entities.ListItem, err error) {
	queryString := `
		SELECT list_item_id, li.name, li.description
		FROM list_item li
		JOIN user u ON u.user_id = li.user_id
		WHERE u.user_uuid = ? ;
	`

	rows, err := r.db.Query(queryString, userUUID)
	if err != nil {
		log.Fatal("erro", err)
	}

	for rows.Next() {
		var ListItem entities.ListItem
		rows.Scan(&ListItem.ListItemId, &ListItem.Name, &ListItem.Description)

		userList = append(userList, ListItem)
	}

	return userList, err
}

func (r listItemRepo) CreateListItem(userId int, listItemUUID string, name string, description string) (err error) {
	queryString := `
	INSERT INTO list_item (user_id, list_item_uuid,  name, description)
	VALUES (?, ?, ?, ?)`

	_, err = r.db.Exec(queryString, userId, listItemUUID, name, description)
	if err != nil {
		return err
	}

	return err
}

func (r listItemRepo) DeleteListItem(uuid string) error {
	queryString := `DELETE FROM list_item
	WHERE list_item_uuid = ?`

	_, err := r.db.Query(queryString, uuid)

	return err
}

func (r listItemRepo) UpdateListItem(name string, description string, uuid string) error {
	queryString := `
	UPDATE list_item
	SET name = ?, description = ?
	WHERE list_item_uuid = ?
	LIMIT 1;`

	_, err := r.db.Query(queryString, name, description, uuid)

	return err
}
