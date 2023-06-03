package server

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

func InitListItemRoutes(e *echo.Echo, db *sql.DB) {
	controller := controller{
		DB: db,
	}

	e.GET("/userList/:user_id", controller.GetUserList)

	e.POST("/createItem/:user_id", controller.CreateUserItem)
}
