package controller

import (
	"net/http"

	"github.com/quniverse26/miniproject/model"
	"github.com/quniverse26/miniproject/config"

	"github.com/labstack/echo"
)

func Register(c echo.Context) error {
	
	reqAuth := model.BuyerRegister{}

	c.Bind(&reqAuth)

	err := config.DB.First(&model.Buyer{},"email = ?", reqAuth.Email).Error
	
	if err == nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": "email already registered",
		})
	}

	newBuyer := model.Buyer{
		Name: reqAuth.Name,
		Email: reqAuth.Email,
		Password: reqAuth.Password,
	}

	if err := config.DB.Create(&newBuyer).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": "failed to create buyer",
		})
	}
	
	return c.JSON(http.StatusCreated, echo.Map{
		"message": "success created buyer",
	})
}

func Login(c echo.Context) error {
	
	reqAuth := model.BuyerLogin{}

	c.Bind(&reqAuth)

	buyer := model.Buyer{}

	err := config.DB.First(&buyer,"email = ?", reqAuth.Email).Error
	
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": "email not registered",
		})
	}

	if !buyer.CheckPassword(reqAuth.Password) {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": "invalid email or password",
		})
	}

	token, err := buyer.GenerateToken()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": "failed to generate token",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success login",
		"token": token,
	})
}