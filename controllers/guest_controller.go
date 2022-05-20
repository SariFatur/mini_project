package controllers

import (
	"myproject/config"
	"myproject/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateGuestController(c echo.Context) error {
	guest := model.Guest{}
	c.Bind(&guest)

	err := config.DB.Create(&guest).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "success create account",
		"guest":   guest,
	})
}

func GetGuestByIdController(c echo.Context) error {
	var guest model.Guest

	stringId := c.Param("id")
	err := config.DB.First(&guest, "id = ?", stringId).Debug().Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get guest by id",
		"guest":   guest,
	})
}

func GetGuestController(c echo.Context) error {
	var guest []model.Guest

	err := config.DB.Find(&guest).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    guest,
	})
}

func DeleteGuestController(c echo.Context) error {
	stringId := c.Param("id")
	err := config.DB.Delete(&model.Guest{}, "id = ?", stringId).Debug().Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete account with idGuest `" + stringId + "`",
	})
}

func UpdateGuestController(c echo.Context) error {
	guest := model.Guest{}
	c.Bind(&guest)
	stringId := c.Param("id")
	err := config.DB.Model(&guest).Where("id = ?", stringId).Updates(guest).Debug().Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"guest":   guest,
	})
}
