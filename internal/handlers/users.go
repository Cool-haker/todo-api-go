package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Cool-haker/todo-api-go/internal/db"
	"github.com/Cool-haker/todo-api-go/internal/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func PostUserHandler(c echo.Context) error {
	var user models.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid JSON",
		})
	}

	if user.Email == "" || user.Username == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "email are required",
		})
	}

	if err := db.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"user": user,
	})
}

func GetUserHandler(c echo.Context) error {
	var user models.User

	userParam := c.Param("id")
	userId, err := strconv.Atoi(userParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid user id",
		})
	}

	if err := db.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "user not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}
