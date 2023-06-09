package server

import (
	"todo-list-api/data"

	"github.com/labstack/echo/v4"
)

func InitUserRoutes(e *echo.Echo, db data.CnnMask) {
	controller := controller{
		DB: db,
	}

	e.POST("/registerUser", controller.HandleRegisterUser)
}
