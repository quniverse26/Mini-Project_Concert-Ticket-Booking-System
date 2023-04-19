package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	id_user int `json:"id_user" form:"id_user"`
	name string `json:"name" form:"name"`
	phone_number string `json:"phone_number" form:"phone_number"`
	email string `json:"email" form:"email"`
	password string `json:"password" form:"password"`
}

var users []User