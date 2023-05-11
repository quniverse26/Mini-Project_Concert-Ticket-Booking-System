package controller

import (
	"net/http"
	//"strconv"

	"github.com/quniverse26/miniproject/model"
	"github.com/quniverse26/miniproject/config"

	"github.com/labstack/echo"
)

// get all admins
func GetAdminsController(c echo.Context) error {
	var admins []model.Admin

	config.DB.Find(&admins)
  
	if err := config.DB.Find(&admins).Error; err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
	  "message": "success get all admins",
	  "admin":   admins,
	})
  }

  // create new admin
  func CreateAdminController(c echo.Context) error {
	admins := model.Admin{}
	c.Bind(&admins)
  
  
	if err := config.DB.Save(&admins).Error; err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
	  "message": "success create new admin",
	  "admin":    admins,
	})
  }