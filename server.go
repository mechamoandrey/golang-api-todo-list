package main

import (
	"log"
	"net/http"
	"todo-list-api/data"
	"todo-list-api/infra/config"
	serverList "todo-list-api/server/list-item"
	serverUser "todo-list-api/server/user"

	"github.com/labstack/echo/v4"
)

func main() {
	config, err := config.GetConfig()
	if err != nil {
		log.Fatal("unable to read config", err)
	}

	e := echo.New()

	cnn, err := data.InitDB(config)
	if err != nil {
		log.Fatal("unable to use db connection", err)
	}

	serverUser.InitUserRoutes(e, cnn)
	serverList.InitListItemRoutes(e, cnn)

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
