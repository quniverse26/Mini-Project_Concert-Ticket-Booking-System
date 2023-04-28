package service

import (
	"fmt"
	"testing"

	"github.com/quniverse26/miniproject/dao"
	"github.com/quniverse26/miniproject/model"
	"github.com/stretchr/testify/assert"
)

type TicketService struct {
	ticketDao dao.TicketDaoInterface
}

func (t TicketService) GetTicket(id int) (model.Ticket, error) {
	return t.ticketDao.GetTicket(id)
}

type TicketDaoMockFoundTicket struct{}

func (t *TicketDaoMockFoundTicket) GetTicket(id int) (model.Ticket, error) {
	return model.Ticket{
		Ticket_name: "Monokrom",
	}, nil
}

type TicketDaoMockNotFoundTicket struct {}

func (t *TicketDaoMockNotFoundTicket) GetTicket(id int) (model.Ticket, error)  {
	return model.Ticket{}, fmt.Errorf("ticket not found")
}

func Test_TicketService_GetTicket(t *testing.T)  {
	type fields struct{
		ticketDao dao.TicketDaoInterface
	}
	type args struct {
		id int
	}
	tests := []struct {
		Ticket_name string
		fields fields
		args args
		want model.Ticket
		wantErr error
	}{
		{
		Ticket_name: `should return ticket when found`,
		fields: fields{
			ticketDao: &TicketDaoMockFoundTicket{},
		},
		args: args{10},
		want: model.Ticket{Ticket_name: "Monokrom"},
	},
	{
		Ticket_name: `should return error when not found`,
		fields: fields{
			ticketDao: &TicketDaoMockNotFoundTicket{},
		},
		args: args{10},
		wantErr: fmt.Errorf("ticket not found"),
	},
}

for _, tt := range tests {
	t.Run(tt.Ticket_name, func(t *testing.T) {
		x := TicketService {
			ticketDao: tt.fields.ticketDao,
		}
		got, err := x.GetTicket(tt.args.id)
		assert.Equal(t, tt.want, got)
		assert.Equal(t, tt.wantErr, err)
	})
}
}