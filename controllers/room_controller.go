package controllers

import (
	"myproject/model"
	"myproject/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllRoomController(c echo.Context) error {
	listRoom, err := repository.GetAllRoomController()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, listRoom)
}

func CreateRoomController(c echo.Context) error {
	id := c.Param("id")
	room := model.Room{}
	c.Bind(&room)
	err := repository.CreateRoomController(room)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success insert room dengan idRoom '" + id + "'",
	})
}

// func CreateRoomController(c echo.Context) error {
// 	room := model.Room{}
// 	c.Bind(&room)

// 	err := config.DB.Save(&room).Error
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 			"message": err.Error(),
// 		})
// 	}
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"massage": "success create room",
// 		"room":    room,
// 	})
// }
