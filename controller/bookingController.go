package controller

import (
	"net/http"
	"strconv"

	"github.com/quniverse26/miniproject/model"
	"github.com/quniverse26/miniproject/config"

	"github.com/labstack/echo"
)

// get all bookings
func GetBookingsController(c echo.Context) error {
	var bookings []model.Booking

	config.DB.Find(&bookings)
  
	if err := config.DB.Find(&bookings).Error; err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
	  "message": "success get all bookings",
	  "booking":   bookings,
	})
  }

// create new booking
func CreateBookingController(c echo.Context) error {
	bookings := model.Booking{}
	c.Bind(&bookings)
  
  
	if err := config.DB.Save(&bookings).Error; err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
	  "message": "success create new booking",
	  "booking":    bookings,
	})
}

//get booking by id
func GetBookingController(c echo.Context) error {
	// Bind request data to buyer struct
	booking := model.Booking{}
	if err := c.Bind(&booking); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Find booking by ID
	id, _ := strconv.Atoi(c.Param("id"))
	if err := config.DB.Find(&booking, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, booking)
}