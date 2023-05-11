package main

import (
	"github.com/quniverse26/miniproject/config"
	"github.com/quniverse26/miniproject/routes"

	//"github.com/jinzhu/gorm"
  	_ "github.com/jinzhu/gorm/dialects/mysql"
  	//"github.com/labstack/echo"
)

func main() {
	config.InitDB()
	e := routes.New()
	
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
  }