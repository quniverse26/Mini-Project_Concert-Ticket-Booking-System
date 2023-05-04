package model

import (

	"github.com/jinzhu/gorm"
)

type Purchasings struct {
	gorm.Model
	Status string `json:"status" form:"status"`
	Buyers []Buyer `json:"buyers" form:"-" gorm:"many2many:purchasings"`
	Tickets []Ticket `json:"tickets" form:"-" gorm:"many2many:purchasings"`
	Bookings []Booking `json:"bookings" form:"-" gorm:"many2many:purchasings"`

	//Status string `json:"status" form:"status"`
	//Buyers []Buyer `json:"buyers" form:"-" gorm:"ForeignKey:id_buyer;association_foreignkey:buyers(id);"`
	//Tickets []Ticket `json:"tickets" form:"-" gorm:"ForeignKey:id_ticket;association_foreignkey:tickets(id);"`
	//Bookings []Booking `json:"bookings" form:"-" gorm:"ForeignKey:id_booking;association_foreignkey:bookings(id);"`
}

var purchasings []Purchasings