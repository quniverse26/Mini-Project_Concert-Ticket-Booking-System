package controller

import (
	"net/http"
	//"strconv"

	"github.com/quniverse26/miniproject/model"
	"github.com/quniverse26/miniproject/config"

	"github.com/labstack/echo"
)

// get all tickets
func GetTicketsController(c echo.Context) error {
	var tickets []model.Ticket

	config.DB.Find(&tickets)
  
	if err := config.DB.Find(&tickets).Error; err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
	  "message": "success get all tickets",
	  "ticket":   tickets,
	})
  }

  // create new ticket
  func CreateTicketController(c echo.Context) error {
	tickets := model.Ticket{}
	c.Bind(&tickets)
  
  
	if err := config.DB.Save(&tickets).Error; err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
	  "message": "success create new ticket",
	  "ticket":    tickets,
	})
  }