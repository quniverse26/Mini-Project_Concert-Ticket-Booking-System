package model

import (
	"github.com/jinzhu/gorm"
)

type Ticket struct {
	gorm.Model
	Ticket_name string `json:"ticket_name" form:"ticket_name"`
	Event_name string `json:"event_name" form:"event_name"`
	Price int `json:"price" form:"price"`
	Qty int `json:"qty" form:"qty"`
	Seat string `json:"seat" form:"seat"`
}

var tickets []Ticket