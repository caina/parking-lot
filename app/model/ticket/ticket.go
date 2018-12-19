package ticket

import (
	"github.com/parking-lot/app/model/car"
	"time"
)

type Ticket struct {
	time time.Time
	code string
	car  car.Car
}

func SueTicket() *Ticket {
	return &Ticket{
		time: time.Now(),
		code: Generate(),
	}
}

func (ticket *Ticket) SetCar(car car.Car) bool {
	if &ticket.car != nil {
		ticket.car = car
		return true
	}
	return false
}

func (ticket *Ticket) ToString() string {
	return ticket.time.Format("3:04 PM") + " > " + ticket.car.ToString() + ", ticket: " + ticket.code
}

func (ticket *Ticket) GetCode() string {
	return ticket.code
}
