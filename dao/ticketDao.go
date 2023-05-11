package dao

import (
	"github.com/quniverse26/miniproject/model"
)

type TicketDaoInterface interface {
	GetTicket(id int) (model.Ticket, error)
}