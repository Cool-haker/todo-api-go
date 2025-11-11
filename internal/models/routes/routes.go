package routes

import (
	hand "github.com/Cool-haker/todo-api-go/internal/handlers"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	api := e.Group("/api/v1")

	todos := api.Group("/todos")
	// todos.Use(middleware.JWTAuth)
	todos.GET("/:id", hand.GetTodoHandler)
	todos.POST("", hand.PostTodoHandler)
	todos.PATCH("/:id", hand.PatchTodoHandler)
	todos.DELETE("/:id", hand.DeleteTodoHandler)

	users := api.Group("/users")
	users.POST("/", hand.RegisterUserHandler)
	users.GET("/:id", hand.LoginUserHandler)

	auth := api.Group("/auth")
	auth.POST("/register", hand.RegisterUserHandler)
	auth.POST("/login", hand.LoginUserHandler)

}
