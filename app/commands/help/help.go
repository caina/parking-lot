package help

import (
	"bytes"
	"github.com/parking-lot/app/config"
)

func Command(client *config.Client) {
	buffer := bytes.Buffer{}
	buffer.WriteString("list of commands available \n")
	buffer.WriteString("-- initialization \n")
	buffer.WriteString("-- status | list of cars parked \n")
	buffer.WriteString("-- park {plate color} | park the car on the slot nearby the door \n")
	buffer.WriteString("-- leave {ticket-id} | remove the car from the park \n")

	client.Send <- buffer.String()
}
