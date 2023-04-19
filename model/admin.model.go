package model

import (
	"github.com/jinzhu/gorm"
)

type Admin struct {
	gorm.Model
	Id_admin int `json:"id_admin" form:"id_admin"`
	Name string `json:"name" form:"name"`
	Phone_number string `json:"phone_number" form:"phone_number"`
	Email string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var admins []Admin