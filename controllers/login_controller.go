package controllers

import (
	"encoding/base64"
	"fmt"
	"myproject/repository"
	"myproject/utils"
	"strings"

	"github.com/labstack/echo/v4"
)

type LoginController struct {
	iUserRepo repository.IUserRepo
}

func NewLoginController(iUserRepo repository.IUserRepo) LoginController {
	return LoginController{iUserRepo}
}

func (ctrl LoginController) Login(c echo.Context) error {
	authorization := c.Request().Header.Get("authorization")
	arrAuth := strings.Split(authorization, " ")
	if len(arrAuth) != 2 {
		fmt.Println("header auth invalid")
		return c.JSON(400, map[string]interface{}{
			"message": "header auth invalid",
		})
	} else if arrAuth[0] != "Basic" {
		fmt.Println("header auth must be basic")
		return c.JSON(400, map[string]interface{}{
			"message": "header auth must be basic",
		})
	}
	// "Basic am9objpzZWNyZXQ="
	var decodeByte, _ = base64.StdEncoding.DecodeString(arrAuth[1])
	// john:secret
	arrDecodeString := strings.Split(string(decodeByte), ":")
	username, password := arrDecodeString[0], arrDecodeString[1]
	user, err := ctrl.iUserRepo.GetUserByUsername(username)
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "user not registered",
		})
	} else if user.Password != password {
		return c.JSON(400, map[string]interface{}{
			"message": "password incorect",
		})
	}
	jwToken := utils.CreateJWT(user)
	return c.JSON(400, map[string]interface{}{
		"token": jwToken,
	})
}
