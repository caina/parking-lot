package park

import (
	"github.com/parking-lot/app"
	"github.com/parking-lot/app/constants"
	"github.com/parking-lot/app/model/car"
	"github.com/parking-lot/app/model/ticket"
)

func Command(args []string, client *app.Client) {
	if len(args) < 2 {
		client.Send <- constants.CommandMalformatted
		return
	}

	newTicket := ticket.SueTicket()
	parkingCar := &car.Car{
		Plate: args[0],
		Color: args[1],
	}
	_ = newTicket.SetCar(*parkingCar)

	if err := client.Parking.Park(*newTicket); err != nil {
		client.Send <- err.Error()
		return
	}

	client.Send <- "Ticket generated: " + newTicket.ToString()
}
