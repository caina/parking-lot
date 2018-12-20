package leave

import (
	. "github.com/parking-lot/app/client"
	"github.com/parking-lot/app/model/ticket"
	"strings"
)

func Command(args []string, client *Client) {
	currentTickets := client.Parking.GetTickets()
	tickets := make([]ticket.Ticket, 0)

	for i := 0; i < len(currentTickets); i++ {
		if strings.Compare(currentTickets[i].GetCode(), args[0]) != 0 {
			tickets = append(tickets, currentTickets[i])
		}
	}

	client.Parking.SetTickets(tickets)
	client.Send <- "Good by!"
}
