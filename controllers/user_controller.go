package controllers

import (
	"myproject/config"
	"myproject/middleware"
	"myproject/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateUserController(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)

	err := config.DB.Save(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "success create user",
		"user":    user,
	})
}

func GetUserController(c echo.Context) error {
	var users []model.User

	err := config.DB.Find(&users).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    users,
	})
}

func LoginUserController(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)

	err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "fail login",
			"error":   err.Error(),
		})
	}

	token, err := middleware.CreateToken(user.ID, user.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "fail login",
			"error":   err.Error(),
		})
	}

	userResponse := model.UserResponse{user.ID, user.Username, user.Email, token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "success create user",
		"user":    userResponse,
	})
}
