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

func PostTodoHandler(c echo.Context) error {
	var todo models.Todo

	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid JSON",
		})
	}

	if todo.UserID == 0 || todo.Title == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "user_id and title  are required",
		})
	}

	if err := db.DB.Create(&todo).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Todo created successfully",
		"todo":    todo,
	})
}

func GetTodoHandler(c echo.Context) error {
	var todos []models.Todo

	userParam := c.Param("id")
	userId, err := strconv.Atoi(userParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid user id",
		})
	}

	if err := db.DB.Where("user_id = ?", userId).Find(&todos).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "todos not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"todos": todos,
	})
}

func PatchTodoHandler(c echo.Context) error {
	var tmpTodo models.TmpTodo
	var todo models.Todo

	idParam := c.Param("id")
	todoID, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid todo id",
		})
	}

	if err := c.Bind(&tmpTodo); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid JSON",
		})
	}

	if err := db.DB.Where("id = ?", todoID).First(&todo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "todos not found",
			})
		}

		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	if tmpTodo.Title != nil {
		todo.Title = *tmpTodo.Title
	}
	if tmpTodo.Description != nil {
		todo.Description = *tmpTodo.Description
	}
	if tmpTodo.IsCompleted != nil {
		todo.IsCompleted = *tmpTodo.IsCompleted
	}

	if err := db.DB.Model(&todo).Updates(todo).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"updatedTodo": todo,
	})
}

func DeleteTodoHandler(c echo.Context) error {
	var todo models.Todo

	idParam := c.Param("id")
	todoID, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid todo id",
		})
	}

	if err := db.DB.Where("id = ?", todoID).First(&todo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "todo not found",
			})
		}

		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	if err := db.DB.Delete(&todo).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "todo deleted successfully",
	})
}
