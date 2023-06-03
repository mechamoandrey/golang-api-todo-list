package main

import (
	"log"
	"net/http"
	"todo-list-api/data"
	serverList "todo-list-api/server/list-item"
	serverUser "todo-list-api/server/user"

	"github.com/labstack/echo/v4"
)

// type user struct {
// 	Id       int
// 	Name     string
// 	userList []userListItem
// }

func main() {

	var err error

	e := echo.New()

	db, err := data.InitDB()
	if err != nil {
		log.Fatal("unable to use db connection", err)
	}

	serverUser.InitUserRoutes(e, db)
	serverList.InitListItemRoutes(e, db)

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
