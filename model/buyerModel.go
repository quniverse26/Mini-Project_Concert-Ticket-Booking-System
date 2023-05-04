package model

import (
	"github.com/jinzhu/gorm"
)

type Buyer struct {
	gorm.Model
	Name string `json:"name" form:"name"`
	Phone_number string `json:"phone_number" form:"phone_number"`
	Email string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var buyers []Buyer