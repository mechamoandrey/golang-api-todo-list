package server

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

func InitUserRoutes(e *echo.Echo, db *sql.DB) {
	controller := controller{
		DB: db,
	}

	e.GET("/user/:user_id", controller.GetUser)
	e.POST("/createUser", controller.CreateUserData)
}
