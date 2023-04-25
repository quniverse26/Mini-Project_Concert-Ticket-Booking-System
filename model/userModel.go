package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
	Phone_number string `json:"phone_number"`
	Email string `json:"email"`
	Password string `json:"password"`
}

var users []User