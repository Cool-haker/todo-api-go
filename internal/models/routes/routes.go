package routes

import (
	"github.com/labstack/echo/v4"
)

func initRoutes(e *echo.Echo) {
	// api := e.Group("/api/v1")

	todos := e.Group("/todos")
	todos.GET("/:id", hand.getTodoHandler)
	todos.POST("", postTodoHandler)
	todos.PATCH("/:id", patchTodoHandler)
	todos.DELETE("/:id", deleteTodoHandler)

	e.GET("/:id", getUserHandler)
	e.POST("", postUserHandler)

	e.Start(":8080")
}
