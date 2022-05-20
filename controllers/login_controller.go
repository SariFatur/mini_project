package controllers

import (
	"myproject/config"
	"myproject/middleware"
	"myproject/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginController(c echo.Context) error {
	guest := model.Guest{}
	c.Bind(&guest)

	err := config.DB.Where("email = ? AND password = ?", guest.Email, guest.Password).First(&guest).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "fail login",
			"error":   err.Error(),
		})
	}

	token, err := middleware.CreateToken(guest.Id, guest.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "fail login",
			"error":   err.Error(),
		})
	}

	guestResponse := model.GuestResponse{guest.Id, guest.Name, guest.Email, token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "success",
		"user":    guestResponse,
	})
}
