package controllers

import (
	"myproject/config"
	"myproject/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreatePaymentMethodController(c echo.Context) error {
	room := model.Payment_method{}
	c.Bind(&room)

	err := config.DB.Save(&room).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "success create payment method",
		"room":    room,
	})
}

func GetPaymentMethodController(c echo.Context) error {
	var PM []model.Payment_method

	err := config.DB.Find(&PM).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    PM,
	})
}

func DeletePaymentMethodController(c echo.Context) error {
	stringId := c.Param("id")
	err := config.DB.Delete(&model.Payment_method{}, "id = ?", stringId).Debug().Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete payment method with id `" + stringId + "`",
	})
}

func UpdatePaymentMethodController(c echo.Context) error {
	PM := model.Payment_method{}
	c.Bind(&PM)
	stringId := c.Param("id")
	err := config.DB.Model(&PM).Where("id = ?", stringId).Updates(PM).Debug().Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"PM":      PM,
	})
}
