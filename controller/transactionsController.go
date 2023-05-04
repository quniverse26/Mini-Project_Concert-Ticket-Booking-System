package controller

import (
	"net/http"
	//"strconv"

	"github.com/quniverse26/miniproject/model"
	"github.com/quniverse26/miniproject/config"

	"github.com/labstack/echo"
)

// get all transactions
func GetTransactionsController(c echo.Context) error {
	var transactions []model.Transactions

	config.DB.Find(&transactions)
	config.DB.Preload("Buyers").Find(&transactions)
	config.DB.Preload("Tickets").Find(&transactions)
	config.DB.Preload("Bookings").Find(&transactions)
  
	if err := config.DB.Find(&transactions).Error; err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
	  "message": "success get all transactions",
	  "transaction":   transactions,
	})
  }

  // create new transaction
  func CreateTransactionController(c echo.Context) error {
	transactions := model.Transactions{}
	c.Bind(&transactions)
  
  
	if err := config.DB.Save(&transactions).Error; err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
	  "message": "success create new transaction",
	  "transaction":    transactions,
	})
  }