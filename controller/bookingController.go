package controller

import (
	"net/http"
	//"strconv"

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

  // create new admin
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