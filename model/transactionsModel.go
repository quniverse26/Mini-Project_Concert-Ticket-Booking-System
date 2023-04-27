package model

import (

	"github.com/jinzhu/gorm"
)

type Transactions struct {
	gorm.Model
	Status string `json:"status" form:"status"`
	Users []User `json:"-" form:"-"`
	Tickets []Ticket `json:"-" form:"-"`
	Bookings []Booking `json:"-" form:"-"`
}

var transactions []Transactions