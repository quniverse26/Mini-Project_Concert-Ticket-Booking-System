package routes

import (
	//"net/http"
	
	"github.com/quniverse26/miniproject/controller"
	//"github.com/quniverse26/miniproject/middleware"

	"github.com/labstack/echo"
	//"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	//e.Pre(middleware.RemoveTrailingSlash())

	// Route / to handler function

	// auth
	authGroup := e.Group("/auth")
	{
		authGroup.POST("/register", controller.Register)
		authGroup.POST("/login", controller.Login)
	}

	//buyer
	e.GET("/buyers", controller.GetBuyersController)
	e.GET("/buyer/:id", controller.GetBuyerController)
	e.POST("/buyer", controller.CreateBuyerController)
	e.DELETE("auth/buyer/:id", controller.DeleteBuyerController)
	e.PUT("auth/buyer/:id", controller.UpdateBuyerController)

	//admin
	e.GET("/admins", controller.GetAdminsController)
	e.POST("/admin", controller.CreateAdminController)

	//ticket
	e.GET("/tickets", controller.GetTicketsController)
	e.GET("/ticket/:id", controller.GetTicketController)
	e.POST("/ticket", controller.CreateTicketController)
	e.DELETE("/ticket/:id", controller.DeleteTicketController)
	e.PUT("/ticket/:id", controller.UpdateTicketController)

	//booking
	e.GET("/bookings", controller.GetBookingsController)
	e.GET("/booking/:id", controller.GetBookingController)
	e.POST("/booking", controller.CreateBookingController)

	//purchasing
	e.GET("/purchasings", controller.GetPurchasingsController)

	return e
}