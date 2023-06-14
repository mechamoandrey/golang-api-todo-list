package server

import (
	"log"
	"net/http"
	"todo-list-api/data"
	"todo-list-api/infra/hashPasswd"

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
	password, _ := hashPasswd.HashPassword(requestBody.Password)

	err := c.DB.UserRepo().CreateUser(uuid, name, login, password)
	if err != nil {
		log.Fatal("erro ao inserir o usuário no banco de dados: ", err)
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (c controller) HandleLogin(ctx echo.Context) error {
	requestBody := new(LoginRequest)
	if err := ctx.Bind(requestBody); err != nil {
		return err
	}

	login := requestBody.Login
	passwrd := requestBody.Password

	user, err := c.DB.UserRepo().GetUserByLogin(login)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	isValid := hashPasswd.CheckPasswordHash(passwrd, string(user.Password))
	if !isValid {
		return ctx.JSON(http.StatusForbidden, "Usuário ou senha incorreta")
	}

	return ctx.NoContent(http.StatusNoContent)
}
