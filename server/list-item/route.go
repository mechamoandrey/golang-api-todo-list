package server

import (
	"todo-list-api/data"

	"github.com/labstack/echo/v4"
)

func InitListItemRoutes(e *echo.Echo, db data.CnnMask) {
	controller := controller{
		DB: db,
	}

	e.GET("/userList/:user_uuid", controller.HandleGetUserList)
	e.POST("/createItem/:user_id", controller.HandleCreateListItem)
	e.DELETE("/deleteItem/:list_item_uuid", controller.HandleDeleteListItem)
	e.PUT("/updateItem/:list_item_uuid", controller.HandleUpdateListItem)
}
