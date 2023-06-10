package server

import (
	"log"
	"net/http"
	"strconv"
	"todo-list-api/data"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type controller struct {
	DB data.CnnMask
}

func (c controller) HandleGetUserList(ctx echo.Context) error {

	idString := ctx.Param("user_id")
	id, _ := strconv.Atoi(idString)

	userList, err := c.DB.ListItemRepo().GetListItemByUserId(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, "Deu pau")
	}

	return ctx.JSON(http.StatusOK, userList)
}

func (c controller) HandleCreateListItem(ctx echo.Context) error {
	idString := ctx.Param("user_id")
	userId, _ := strconv.Atoi(idString)

	requestBody := new(CreateListItemRequest)
	if err := ctx.Bind(requestBody); err != nil {
		return err
	}

	name := requestBody.Name
	description := requestBody.Description
	uuid := uuid.New().String()

	err := c.DB.ListItemRepo().CreateListItem(userId, uuid, name, description)
	if err != nil {
		log.Fatal("erro ao inserir o item no banco de dados: ", err)
	}

	response := CreateListItemResponse{
		ListItemUUID: uuid,
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
