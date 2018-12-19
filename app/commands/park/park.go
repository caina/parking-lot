package park

import (
	"github.com/parking-lot/app/config"
	"github.com/parking-lot/app/model/car"
	"github.com/parking-lot/app/model/ticket"
)

func Command(args []string, client *config.Client) {
	newTicket := ticket.SueTicket()
	parkingCar := &car.Car{
		Plate: args[0],
		Color: args[1],
	}
	newTicket.SetCar(*parkingCar)

	if parked := client.Parking.Park(*newTicket); parked {
		client.Send <- "Ticket generated: " + newTicket.ToString()
	} else {
		client.Send <- "FAIL: Not able to park the car, the parking lot is full!"
	}

}
