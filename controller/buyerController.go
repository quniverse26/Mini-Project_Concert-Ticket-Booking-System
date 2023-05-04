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

// get all buyers
func GetBuyersController(c echo.Context) error {
	var buyers []model.Buyer

	config.DB.Find(&buyers)
  
	if err := config.DB.Find(&buyers).Error; err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
	  "message": "success get all buyers",
	  "buyer":   buyers,
	})
  }

  // create new admin
  func CreateBuyerController(c echo.Context) error {
	buyers := model.Buyer{}
	c.Bind(&buyers)
  
  
	if err := config.DB.Save(&buyers).Error; err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
	  "message": "success create new buyer",
	  "buyer":    buyers,
	})
  }

//delete ticket
func DeleteBuyerController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid buyer ID"})
	}

	// check if buyer exists
	buyer := model.Buyer{}
	if err := config.DB.First(&buyer, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, echo.Map{"message": "Buyer not found"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to retrieve buyer"})
	}

	// delete buyer
	if err := config.DB.Unscoped().Delete(&buyer).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to delete buyer"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Buyer deleted successfully"})
}

//update buyer
func UpdateBuyerController(c echo.Context) error {
	data := echo.Map{
		"message": "success update buyer",
	}

	var buyer model.Buyer
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		data["message"] = err.Error()
		return c.JSON(http.StatusBadRequest, data)
	}

	// load buyer from database
	if err := config.DB.First(&buyer, id).Error; err != nil {
		data["message"] = err.Error()
		return c.JSON(http.StatusBadRequest, data)
	}

	// bind updated data to buyer
	if err := c.Bind(&buyer); err != nil {
		data["message"] = err.Error()
		return c.JSON(http.StatusBadRequest, data)
	}

	// update buyer
	if err := config.DB.Save(&buyer).Error; err != nil {
		data["message"] = err.Error()
		return c.JSON(http.StatusBadRequest, data)
	}

	return c.JSON(http.StatusOK, data)
}