package controllers

import (
	"myproject/config"
	"myproject/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateReservationController(c echo.Context) error {
	reservation := model.Reservation{}
	c.Bind(&reservation)

	err := config.DB.Create(&reservation).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage":     "success create reservation",
		"reservation": reservation,
	})
}

func GetReservationController(c echo.Context) error {
	var reservation []model.Reservation

	err := config.DB.Find(&reservation).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    reservation,
	})
}

func DeleteReservationController(c echo.Context) error {
	stringId := c.Param("id")
	err := config.DB.Delete(&model.Reservation{}, "id = ?", stringId).Debug().Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete reservation with id `" + stringId + "`",
	})
}

func UpdateReservationController(c echo.Context) error {
	reservation := model.Reservation{}
	c.Bind(&reservation)
	stringId := c.Param("id")
	err := config.DB.Model(&reservation).Where("id = ?", stringId).Updates(reservation).Debug().Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":     "success",
		"reservation": reservation,
	})
}
