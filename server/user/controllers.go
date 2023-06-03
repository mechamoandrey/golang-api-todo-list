package server

import (
	"database/sql"
	"fmt"
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

type RequestUserData struct {
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type User struct {
	ID    int
	Name  string
	Login string
}

type controller struct {
	DB *sql.DB
}

func (c controller) GetUser(ctx echo.Context) error {
	queryString := `
		SELECT name, login FROM user
		WHERE user_id = ?;
	`

	idString := ctx.Param("user_id")

	id, _ := strconv.Atoi(idString)

	rows, err := c.DB.Query(queryString, id)
	if err != nil {
		log.Fatal("erro", err)
	}

	var user User

	if rows.Next() {
		err := rows.Scan(&user.Name, &user.Login)
		if err != nil {
			log.Fatal("erro ao ler dados do banco de dados", err)
		}
	} else {
		return ctx.String(http.StatusNotFound, "Usuário não encontrado")
	}

	response := fmt.Sprintf("Usuário: %s, Login: %s", user.Name, user.Login)
	return ctx.String(http.StatusOK, response)
}

func (c controller) CreateUserData(ctx echo.Context) error {

	// Criando uma instância do objeto RequestBody
	requestBody := new(RequestUserData)

	// Realizando o binding do corpo da requisição para o objeto RequestBody
	if err := ctx.Bind(requestBody); err != nil {
		return err
	}

	// Acessando o campo Name do objeto RequestBody
	name := requestBody.Name
	login := requestBody.Login
	password := requestBody.Password

	queryString := `
		INSERT INTO user (name, login, password) 
		VALUES (?, ?, ?);
	`

	_, err := c.DB.Query(queryString, name, login, password)
	if err != nil {
		log.Fatal("erro ao inserir usuário no banco de dados:", err)
	}

	// Exemplo de resposta
	response := struct {
		Message string `json:"message"`
	}{
		Message: fmt.Sprintf("Requisição POST recebida com sucesso."),
	}

	return ctx.JSON(http.StatusOK, response)
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

	return ctx.String(http.StatusOK, "usuário: ")
}
