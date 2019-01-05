package park

import (
	"github.com/parking-lot/app/constants"
	"github.com/parking-lot/app/model/car"
	"github.com/parking-lot/app/model/parking-lot"
	. "github.com/parking-lot/app/model/ticket"
	"github.com/pkg/errors"
)

func Command(args []string, lot *parking_lot.ParkingLot) (*Ticket, error) {
	if len(args) < 2 {
		return nil, errors.New(constants.CommandMalformatted)
	}

	var position = 1
	for position < len(lot.GetTickets())+1 {
		notFound := true
		for _, myTickets := range lot.GetTickets() {
			if myTickets.GetPosition() == position {
				notFound = false
				break
			}
		}

		if notFound {
			break
		}
		position++
	}

	newTicket := SueTicket(position)
	parkingCar := &car.Car{
		Plate: args[0],
		Color: args[1],
	}
	_ = newTicket.SetCar(*parkingCar)

	return newTicket, nil
}
