package controller

import (
	"net/http"
	"strconv"
	"errors"

	"github.com/quniverse26/miniproject/model"
	"github.com/quniverse26/miniproject/config"

	"github.com/labstack/echo"
	"github.com/jinzhu/gorm"
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

//get ticket by id
func GetTicketController(c echo.Context) error {
	// Bind request data to buyer struct
	ticket := model.Ticket{}
	if err := c.Bind(&ticket); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Find ticket by ID
	id, _ := strconv.Atoi(c.Param("id"))
	if err := config.DB.Find(&ticket, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, ticket)
}

//delete ticket
func DeleteTicketController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid ticket ID"})
	}

	// check if ticket exists
	ticket := model.Ticket{}
	if err := config.DB.First(&ticket, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, echo.Map{"message": "Ticket not found"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to retrieve ticket"})
	}

	// delete ticket
	if err := config.DB.Unscoped().Delete(&ticket).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to delete ticket"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Ticket deleted successfully"})
}

//update ticket
func UpdateTicketController(c echo.Context) error {
	data := echo.Map{
		"message": "success update ticket",
	}

	var ticket model.Ticket
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		data["message"] = err.Error()
		return c.JSON(http.StatusBadRequest, data)
	}

	// load user from database
	if err := config.DB.First(&ticket, id).Error; err != nil {
		data["message"] = err.Error()
		return c.JSON(http.StatusBadRequest, data)
	}

	// bind updated data to user
	if err := c.Bind(&ticket); err != nil {
		data["message"] = err.Error()
		return c.JSON(http.StatusBadRequest, data)
	}

	// update user
	if err := config.DB.Save(&ticket).Error; err != nil {
		data["message"] = err.Error()
		return c.JSON(http.StatusBadRequest, data)
	}

	return c.JSON(http.StatusOK, data)
}