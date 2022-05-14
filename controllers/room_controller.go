package controllers

import (
	"myproject/config"
	"myproject/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateRoomController(c echo.Context) error {
	room := model.Room{}
	c.Bind(&room)

	err := config.DB.Save(&room).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "success create room",
		"room":    room,
	})
}

func GetRoomController(c echo.Context) error {
	var room []model.Room

	err := config.DB.Find(&room).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    room,
	})
}

func DeleteRoomController(c echo.Context) error {
	stringId := c.Param("id")
	err := config.DB.Delete(&model.Room{}, "id = ?", stringId).Debug().Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete account with id `" + stringId + "`",
	})
}

func UpdateRoomController(c echo.Context) error {
	room := model.Room{}
	c.Bind(&room)
	stringId := c.Param("id")
	err := config.DB.Model(&room).Where("id = ?", stringId).Updates(room).Debug().Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"room":    room,
	})
}
