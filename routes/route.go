package routes

import (
	"myproject/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	// route user
	// e.GET("/users", controllers.GetUserController)
	// e.POST("/users", controllers.CreateUserController)
	// e.POST("/login", controllers.LoginUserController)

	// eAuthBasic := e.Group("/auth")
	// eAuthBasic.Use(mid.BasicAuth(middleware.BasicAuthDB))
	// eAuthBasic.GET("/users", controllers.GetUserController)

	// eJwt := e.Group("/jwt")
	// eJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	// eJwt.GET("/users", controllers.GetUserController)

	// route guest
	e.GET("/guest", controllers.GetGuestController)
	// m.LogMiddleware(e)
	e.POST("/guest", controllers.CreateGuestController)
	e.DELETE("/guest/:id", controllers.DeleteGuestController)
	e.PUT("/guest/:id", controllers.UpdateGuestController)

	// route room
	e.GET("/room", controllers.GetRoomController)
	e.POST("/room", controllers.CreateRoomController)
	e.DELETE("/room/:id", controllers.DeleteRoomController)
	e.PUT("/room/:id", controllers.UpdateRoomController)

	// route payment method
	e.GET("/PM", controllers.GetPaymentMethodController)
	e.POST("/PM", controllers.CreatePaymentMethodController)
	e.DELETE("/PM/:id", controllers.DeletePaymentMethodController)
	e.PUT("/PM/:id", controllers.UpdatePaymentMethodController)

	// route reservation
	e.GET("/reservation", controllers.GetReservationController)
	e.POST("/reservation", controllers.CreateReservationController)
	e.DELETE("/reservation/:id", controllers.DeleteReservationController)
	e.PUT("/reservation/:id", controllers.UpdateReservationController)

	return e
}
