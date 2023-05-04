package controller

import (
	"net/http"
	//"strconv"
	//"errors"

	"github.com/quniverse26/miniproject/model"
	"github.com/quniverse26/miniproject/config"

	"github.com/labstack/echo"
	//"github.com/jinzhu/gorm"
)

// get all purchasings
// func GetPurchasingsController(c echo.Context) error {
// 	var purchasings []model.Purchasings
// 	config.DB.Find(&purchasings)

// 	err := config.DB.Preload("Buyers").Preload("Tickets").Preload("Bookings").Find(&purchasings).Error
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message":   "success get all purchasings",
// 		"purchasings": purchasings,
// 	})
// }


func GetPurchasingsController(c echo.Context) error {
	var purchasings []model.Purchasings

	config.DB.Find(&purchasings)
	config.DB.Preload("Buyers").Find(&purchasings)
	config.DB.Preload("Tickets").Find(&purchasings)
	config.DB.Preload("Bookings").Find(&purchasings)
  
	if err := config.DB.Find(&purchasings).Error; err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
	  "message": "success get all purchasings",
	  "purchasing":   purchasings,
	})
  }