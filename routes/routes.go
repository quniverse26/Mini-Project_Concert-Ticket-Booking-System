package routes

import (
	"github.com/quniverse26/miniproject/controller"
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	// Route / to handler function
	//user
	e.GET("/user", controller.GetUsersController)
	//e.GET("/user/:id", controller.GetUserController)
	e.POST("/user", controller.CreateUserController)
	e.DELETE("/user/:id", controller.DeleteUserController)
	e.PUT("/user/:id", controller.UpdateUserController)

	return e
}