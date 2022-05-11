package routes

import (
	"myproject/constants"
	"myproject/controllers"
	"myproject/middleware"
	m "myproject/middleware"

	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/users", controllers.GetUserController)
	m.LogMiddleware(e)
	e.POST("/users", controllers.CreateUserController)
	e.POST("/login", controllers.LoginUserController)

	eAuthBasic := e.Group("/auth")
	eAuthBasic.Use(mid.BasicAuth(middleware.BasicAuthDB))
	eAuthBasic.GET("/users", controllers.GetUserController)

	eJwt := e.Group("/jwt")
	eJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eJwt.GET("/users", controllers.GetUserController)

	return e
}
