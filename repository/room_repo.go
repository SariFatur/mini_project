package repository

import (
	"fmt"
	"myproject/config"
	"myproject/model"
)

func GetAllRoomController() ([]model.Room, error) {
	var listRoom []model.Room

	err := config.DB.Find(&listRoom).Error
	if err != nil {
		fmt.Println(err)
	}
	return listRoom, err
}

func GetRoomController(idRoom string) (model.Room, error) {
	var room model.Room
	err := config.DB.First(&room, "idRoom = ?", idRoom).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return room, err
}

func DeleteRoomController(idRoom model.Room) error {
	err := config.DB.Delete(&model.Room{}, "idRoom = ?", idRoom).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func CreateRoomController(room model.Room) error {
	err := config.DB.Save(&room).Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}
