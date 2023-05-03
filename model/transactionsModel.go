package model

import (

	"github.com/jinzhu/gorm"
)

type Transactions struct {
	gorm.Model
	//Status string `json:"status" form:"status"`
	//Users []User `json:"users" form:"-" gorm:"many2many:users_transactions"`
	//Tickets []Ticket `json:"tickets" form:"-" gorm:"many2many:tickets_transactions"`
	//Bookings []Booking `json:"bookings" form:"-" gorm:"many2many:bookings_transactions"`
	
	Status string `json:"status" form:"status"`
	Users []User `json:"users" form:"-" gorm:"ForeignKey:id_user;association_foreignkey:users(ID);"`
	Tickets []Ticket `json:"tickets" form:"-" gorm:"ForeignKey:id_ticket;association_foreignkey:tickets(id);"`
	Bookings []Booking `json:"bookings" form:"-" gorm:"ForeignKey:id_booking;association_foreignkey:bookings(id);"`
}

var transactions []Transactions