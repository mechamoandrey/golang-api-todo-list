package server

import (
	"log"
	"net/http"
	"todo-list-api/data"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type controller struct {
	DB data.CnnMask
}

func (c controller) HandleRegisterUser(ctx echo.Context) error {
	requestBody := new(CreateUserRequest)
	if err := ctx.Bind(requestBody); err != nil {
		return err
	}

	uuid := uuid.New().String()
	name := requestBody.Name
	login := requestBody.Login
	password := requestBody.Password

	err := c.DB.UserRepo().HandleRegisterUser(uuid, name, login, password)
	if err != nil {
		log.Fatal("erro ao inserir o usuário no banco de dados: ", err)
	}

	return ctx.NoContent(http.StatusNoContent)
}
