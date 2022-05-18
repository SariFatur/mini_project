package routes

import (
	"myproject/constants"
	"myproject/controllers"

	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	// route guest
	e.GET("/guest", controllers.GetGuestController)
	e.DELETE("/guest/:id", controllers.DeleteGuestController)
	e.PUT("/guest/:id", controllers.UpdateGuestController)

	// e.POST("/login", controllers.LoginGuestController)
	e.POST("/guest", controllers.CreateGuestController)
	e.POST("/login", controllers.LoginController)

	eJwt := e.Group("/jwt")
	eJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eJwt.GET("/users", controllers.GetGuestController)

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

	// route transactions
	e.GET("/transactions", controllers.GetTransactionsController)
	e.POST("/transactions", controllers.CreateTransactionsController)
	e.DELETE("/transactions/:id", controllers.DeleteTransactionsonstroller)
	e.PUT("/transactions/:id", controllers.UpdateTransactionsController)

	return e
}
