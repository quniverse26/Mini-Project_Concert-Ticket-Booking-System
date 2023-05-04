package routes

import (
	//"net/http"
	
	"github.com/quniverse26/miniproject/controller"
	//"github.com/quniverse26/miniproject/middleware"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	// Route / to handler function

	// protected route
	//e.GET("/protected", func(c echo.Context) error {
		//return c.String(http.StatusOK, "This is a protected route")
	//}, middleware.jwtMiddleware)

	//buyer
	e.GET("/buyers", controller.GetBuyersController)
	e.POST("/buyer", controller.CreateBuyerController)
	e.DELETE("/buyer/:id", controller.DeleteBuyerController)
	e.PUT("/buyer/:id", controller.UpdateBuyerController)

	//admin
	e.GET("/admin", controller.GetAdminsController)
	e.POST("/admin", controller.CreateAdminController)

	//ticket
	e.GET("/ticket", controller.GetTicketsController)
	e.POST("/ticket", controller.CreateTicketController)
	e.DELETE("/ticket/:id", controller.DeleteTicketController)
	e.PUT("/ticket/:id", controller.UpdateTicketController)

	//booking
	e.GET("/booking", controller.GetBookingsController)
	e.POST("/booking", controller.CreateBookingController)

	//purchasing
	e.GET("/purchasings", controller.GetPurchasingsController)

	return e
}