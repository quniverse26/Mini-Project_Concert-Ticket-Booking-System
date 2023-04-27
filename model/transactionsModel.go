package model

import (

	"github.com/jinzhu/gorm"
)

type Transactions struct {
	gorm.Model
	Status string `json:"status" form:"status"`
	Users []User `json:"-" form:"-" gorm:"foreignKey:id"`
	Tickets []Ticket `json:"-" form:"-" gorm:"foreignKey:id"`
	Bookings []Booking `json:"-" form:"-" gorm:"foreignKey:id"`
}

var transactions []Transactions