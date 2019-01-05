package ticket

import (
	"bytes"
	"errors"
	"github.com/parking-lot/app/constants"
	"github.com/parking-lot/app/model/car"
	"strconv"
	"time"
)

type Ticket struct {
	time     time.Time
	position int
	code     string
	car      car.Car
}

func SueTicket(position int) *Ticket {
	return &Ticket{
		time:     time.Now(),
		code:     Generate(),
		position: position,
	}
}

func (ticket *Ticket) SetCar(car car.Car) error {
	if &ticket.car != nil {
		ticket.car = car
		return nil
	}
	return errors.New(constants.CarAlreadySetToTicket)
}

func (ticket *Ticket) GetCar() car.Car {
	return ticket.car
}

func (ticket *Ticket) ToString() string {
	builder := bytes.Buffer{}
	builder.WriteString(ticket.time.Format("3:04 PM"))
	builder.WriteString(" | " + strconv.Itoa(ticket.position) + " | ")
	builder.WriteString(" > " + ticket.car.ToString())
	builder.WriteString(", ticket: " + ticket.code)

	return builder.String()
}

func (ticket *Ticket) GetCode() string {
	return ticket.code
}

func (ticket *Ticket) GetPosition() int {
	return ticket.position
}
