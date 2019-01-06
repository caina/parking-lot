package leave

import (
	. "github.com/parking-lot/app/model/parking-lot"
	"github.com/parking-lot/app/model/ticket"
	"strings"
)

func Command(args []string, lot ParkingLot) ([]ticket.Ticket, error) {
	currentTickets := lot.GetTickets()

	tickets := make([]ticket.Ticket, 0)
	for i := 0; i < len(currentTickets); i++ {
		if strings.Compare(currentTickets[i].GetCode(), args[0]) != 0 {
			tickets = append(tickets, currentTickets[i])
		}
	}

	return tickets, nil
}
