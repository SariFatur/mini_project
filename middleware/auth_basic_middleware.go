package middleware

import (
	"myproject/config"
	"myproject/model"

	"github.com/labstack/echo/v4"
)

func BasicAuthDB(username, password string, c echo.Context) (bool, error) {
	var guest model.Guest
	err := config.DB.Where("email = ? AND password = ?", username, password).First(&guest).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
