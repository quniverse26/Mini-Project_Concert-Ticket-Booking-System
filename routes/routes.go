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

	//user
	e.GET("/user", controller.GetUsersController)
	//e.GET("/user/:id", controller.GetUserController)
	e.POST("/user", controller.CreateUserController)
	e.DELETE("/user/:id", controller.DeleteUserController)
	e.PUT("/user/:id", controller.UpdateUserController)

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

	//transactions
	e.GET("/transactions", controller.GetTransactionsController)
	e.POST("/transactions", controller.CreateTransactionController)

	return e
}