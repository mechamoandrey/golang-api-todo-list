package server

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type userListItem struct {
	Name        string
	Description string
	ListItemId  int
}

type RequestListData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type controller struct {
	DB *sql.DB
}

func (c controller) GetUserList(ctx echo.Context) error {
	userList := []userListItem{}

	queryString := `
		SELECT list_item_id, name, description
		FROM list_item
		WHERE user_id = ?;
	`

	idString := ctx.Param("user_id")
	id, _ := strconv.Atoi(idString)

	rows, err := c.DB.Query(queryString, id)
	if err != nil {
		log.Fatal("erro", err)
	}

	for rows.Next() {
		var userListItem userListItem
		rows.Scan(&userListItem.ListItemId, &userListItem.Name, &userListItem.Description)

		userList = append(userList, userListItem)
	}

	return ctx.JSON(http.StatusOK, userList)
}

func (c controller) CreateUserItem(ctx echo.Context) error {
	queryString := `INSERT INTO list_item (user_id, name, description)
	VALUES (?, ?, ?)`

	idString := ctx.Param("user_id") // Obtém o valor do parâmetro "id" da rota

	id, _ := strconv.Atoi(idString)

	// Criando uma instância do objeto RequestBody
	requestBody := new(RequestListData)
	// Realizando o binding do corpo da requisição para o objeto RequestBody
	if err := ctx.Bind(requestBody); err != nil {
		return err
	}

	name := requestBody.Name
	description := requestBody.Description

	_, err := c.DB.Query(queryString, id, name, description)

	if err != nil {
		log.Fatal("erro ao inserir usuário no banco de dados:", err)
	}

	// Faça o que você precisa fazer com o parâmetro aqui, como recuperar o usuário do banco de dados com o ID fornecido

	// Retorna uma resposta de exemplo com o ID do usuário
	return ctx.String(http.StatusOK, "Inseridinho br kkk usuário: "+idString)
}
