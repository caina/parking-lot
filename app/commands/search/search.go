package search

import (
	"github.com/parking-lot/app/constants"
	"github.com/parking-lot/app/model/parking-lot"
	. "github.com/parking-lot/app/model/ticket"
	"github.com/pkg/errors"
	"strings"
)

func Command(args []string, lot *parking_lot.ParkingLot) ([]Ticket, error) {
	if len(args) != 1 {
		return nil, errors.New(constants.CommandMalformatted)
	}

	color := args[0]
	var ticketsFound []Ticket
	for _, ticket := range lot.GetTickets() {
		if strings.Compare(ticket.GetCar().Color, color) == 0 {
			ticketsFound = append(ticketsFound, ticket)
		}
	}

	if len(ticketsFound) == 0 {
		return nil, errors.New(constants.NoCarFound)
	}

	return ticketsFound, nil
}
