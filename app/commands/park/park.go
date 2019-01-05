package park

import (
	"github.com/parking-lot/app/constants"
	"github.com/parking-lot/app/model/car"
	"github.com/parking-lot/app/model/parking-lot"
	"github.com/parking-lot/app/model/ticket"
	"github.com/pkg/errors"
	"strings"
)

func Command(args []string, lot *parking_lot.ParkingLot) (*string, error) {
	if len(args) < 2 {
		return nil, errors.New(constants.CommandMalformatted)
	}

	newTicket := ticket.SueTicket()
	parkingCar := &car.Car{
		Plate: args[0],
		Color: args[1],
	}
	_ = newTicket.SetCar(*parkingCar)

	if err := lot.Park(*newTicket); err != nil {
		return nil, err
	}

	output := strings.Replace(constants.TicketGenerated, "%c", newTicket.ToString(), 1)
	return &output, nil
}
