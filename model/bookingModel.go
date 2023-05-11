package model

import (
	"github.com/jinzhu/gorm"
)

type Booking struct {
	gorm.Model
	Payment_method string `json:"payment_method" form:"payment_method"`
	Total_qty int `json:"total_qty" form:"total_qty"`
	Total_price int `json:"total_price" form:"total_price"`
}

var bookings []Booking