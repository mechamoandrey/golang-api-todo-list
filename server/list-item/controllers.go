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

func (c controller) HandleGetUserList(ctx echo.Context) error {

	userUUID := ctx.Param("user_uuid")

	userList, err := c.DB.ListItemRepo().GetListItemByUserId(userUUID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, "Deu pau")
	}

	return ctx.JSON(http.StatusOK, userList)
}

func (c controller) HandleCreateListItem(ctx echo.Context) error {
	userUUID := ctx.Param("user_uuid")

	user, err := c.DB.UserRepo().GetUserByUUID(userUUID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	requestBody := new(CreateListItemRequest)
	if err := ctx.Bind(requestBody); err != nil {
		return err
	}

	name := requestBody.Name
	description := requestBody.Description
	listItemUUID := uuid.New().String()

	err = c.DB.ListItemRepo().CreateListItem(user.UserId, listItemUUID, name, description)
	if err != nil {
		log.Fatal("erro ao inserir o item no banco de dados: ", err)
	}

	response := CreateListItemResponse{
		ListItemUUID: listItemUUID,
		Name:         name,
		Description:  description,
	}

	return ctx.JSON(http.StatusOK, response)
}

func (c controller) HandleDeleteListItem(ctx echo.Context) error {
	uuid := ctx.Param("list_item_uuid")

	err := c.DB.ListItemRepo().DeleteListItem(uuid)
	if err != nil {
		log.Fatal("Erro ao deletar item no BD", err)
	}

	return ctx.String(http.StatusOK, "Item Removido")
}

func (c controller) HandleUpdateListItem(ctx echo.Context) error {
	uuid := ctx.Param("list_item_uuid")
	requestBody := new(UpdateListItemRequest)
	if err := ctx.Bind(requestBody); err != nil {
		return err
	}

	name := requestBody.Name
	description := requestBody.Description

	err := c.DB.ListItemRepo().UpdateListItem(name, description, uuid)
	if err != nil {
		log.Fatal("Erro ao atualizar item no BD", err)
	}

	return ctx.String(http.StatusOK, "Item Alterado")
}
