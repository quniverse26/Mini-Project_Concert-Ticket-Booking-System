package model

import (

	"github.com/jinzhu/gorm"
)

type Purchasings struct {
	gorm.Model
	// Status string `json:"status" form:"status"`
	// Buyers []Buyer `json:"buyers" form:"-" gorm:"many2many:purchasings"`
	// Tickets []Ticket `json:"tickets" form:"-" gorm:"many2many:purchasings"`
	// Bookings []Booking `json:"bookings" form:"-" gorm:"many2many:purchasings"`

	Status string `json:"status" form:"status"`
	IdBuyer int `json:"id_buyer" gorm:"column:id_buyer"`
	//Buyers []Buyer `json:"buyers" gorm:"foreignKey:IdBuyer;association_foreignkey:buyers(id);"`
	Buyers []Buyer `json:"buyers" gorm:"foreignKey:IdBuyer"`
	IdTicket int `json:"id_ticket" gorm:"column:id_ticket"`
	// Tickets []Ticket `json:"tickets" gorm:"foreignKey:IdTicket;association_foreignkey:tickets(id);"`
	Tickets []Ticket `json:"tickets" gorm:"foreignKey:IdTicket"`
	IdBooking int `json:"id_booking" gorm:"column:id_booking"`
	// Bookings []Booking `json:"bookings" gorm:"foreignKey:IdBooking;association_foreignkey:bookings(id);"`
	Bookings []Booking `json:"bookings" gorm:"foreignKey:IdBooking"`
}

var purchasings []Purchasings