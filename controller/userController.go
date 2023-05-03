package controller

import (
	"net/http"
	"strconv"

	"github.com/quniverse26/miniproject/model"
	"github.com/quniverse26/miniproject/config"

	"github.com/labstack/echo"
)

// get all users
func GetUsersController(c echo.Context) error {
	var users []model.User

	config.DB.Find(&users)
  
	if err := config.DB.Find(&users).Error; err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
	  "message": "success get all users",
	  "users":   users,
	})
  }

  // create new user
  func CreateUserController(c echo.Context) error {
	users := model.User{}
	c.Bind(&users)
  
  
	if err := config.DB.Save(&users).Error; err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
	  "message": "success create new user",
	  "user":    users,
	})
  }

  //delete user
  func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	// load user from database
	user := model.User{}
	if err := config.DB.First(&user, id).Error; err != nil {
		return err
	}

	// delete user
	if err := config.DB.Unscoped().Delete(&user).Error; err != nil {
		return err
	}

	data := &echo.Map{
		"message": "success",
	}
	return c.JSON(http.StatusOK, data)
}

//update user
func UpdateUserController(c echo.Context) error {
	data := echo.Map{
		"message": "success",
	}

	var user model.User
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		data["message"] = err.Error()
		return c.JSON(http.StatusBadRequest, data)
	}

	// load user from database
	if err := config.DB.First(&user, id).Error; err != nil {
		data["message"] = err.Error()
		return c.JSON(http.StatusBadRequest, data)
	}

	// bind updated data to user
	if err := c.Bind(&user); err != nil {
		data["message"] = err.Error()
		return c.JSON(http.StatusBadRequest, data)
	}

	// update user
	if err := config.DB.Save(&user).Error; err != nil {
		data["message"] = err.Error()
		return c.JSON(http.StatusBadRequest, data)
	}

	return c.JSON(http.StatusOK, data)
}
