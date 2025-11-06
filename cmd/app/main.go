package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	InitDB()
	e.GET("/", getHandler)
	e.POST("/", postHandler)
	e.Start(":8080")
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func getHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, world!")
}

func postHandler(c echo.Context) error {
	var u User
	if err := c.Bind(&u); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}
	return c.JSON(http.StatusCreated, u)
}
