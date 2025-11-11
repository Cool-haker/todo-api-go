package handlers

import (
	"errors"
	"net/http"

	"github.com/Cool-haker/todo-api-go/internal/db"
	utils "github.com/Cool-haker/todo-api-go/internal/middleware"
	"github.com/Cool-haker/todo-api-go/internal/models"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RegisterUserHandler(c echo.Context) error {
	var user models.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid JSON",
		})
	}

	if user.Email == "" || user.Username == "" || len(user.Password) < 8 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "All fields must be filled in",
		})
	}

	hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	user.Password = string(hashPass)

	if err := db.DB.Create(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrCheckConstraintViolated) {
			return c.JSON(http.StatusConflict, map[string]string{
				"error": "username already exists",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	token, err := utils.GenerateJWT(uint(user.ID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"user":  user,
		"token": token,
	})
}

func LoginUserHandler(c echo.Context) error {
	var input models.User
	var user models.User

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid JSON",
		})
	}

	if err := db.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "invalid credentional",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "invalid credentional",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	token, err := utils.GenerateJWT(uint(user.ID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"user":  user,
		"token": token,
	})
}
