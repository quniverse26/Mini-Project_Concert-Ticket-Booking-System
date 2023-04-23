package controller

import (
	"net/http"
	"strconv"

	"github.com/quniverse26/miniproject/model"
	"github.com/quniverse26/miniproject/config"

	"github.com/labstack/echo"
)

type Controller struct {
}

// get all users
func GetUsersController(c echo.Context) error {
	var users []model.User

	config.DB.Find(&users)

	//return c.JSON(http.StatusOK, users)
  
	if err := config.DB.Find(&users).Error; err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
	  "message": "success get all users",
	  "user":   users,
	})
  }

  // get user by id
//func GetUserController(c echo.Context) error {
	//user := model.User{}
	//c.Bind(&user)
  
	//if err := config.DB.Find(&user).Error; err != nil {
	  //return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	//}
	//id, _ := strconv.Atoi(c.Param("id"))
	//users := model.User{Id: id, Name: "name", Email: "email", Password: "password"}
	//return c.JSON(http.StatusOK, users)
  //}

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
	id, _ := strconv.Atoi(c.Param("ID"))

	// binding data
	user := model.User{}
	c.Bind(&user)
	
	//var users model.User
	id, err := strconv.Atoi(c.Param("ID"))
	if err != nil {
		return err
	}

	config.DB.Unscoped().Delete(&user, id)

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
	id, err := strconv.Atoi(c.Param("ID"))
	if err != nil {
		data["message"] = err.Error()
		return c.JSON(http.StatusBadRequest, data)
	}

	name := c.Request().PostFormValue("name")

	result := config.DB.Model(&user).Where("ID = ?", id).Update("name", name)
	if result.Error != nil {
		data["message"] = result.Error
		return c.JSON(http.StatusBadRequest, data)
	}

	return c.JSON(http.StatusOK, data)
}