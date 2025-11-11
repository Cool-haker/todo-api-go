package main

import (
	"github.com/Cool-haker/todo-api-go/internal/db"
	"github.com/Cool-haker/todo-api-go/internal/middleware"
	"github.com/Cool-haker/todo-api-go/internal/models/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	db.InitDB()
	middleware.LoadJWTSecret()

	e := echo.New()
	routes.InitRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
