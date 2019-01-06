package help

import (
	"bytes"
)

func Command() string {
	buffer := bytes.Buffer{}
	buffer.WriteString("list of commands available \n")
	buffer.WriteString("-- init {number of slots} \n")
	buffer.WriteString("-- status | list of cars parked \n")
	buffer.WriteString("---- {time} {park_place} {car information} {ticket} \n")
	buffer.WriteString("-- park {plate} {color} | park the car if we have any space left \n")
	buffer.WriteString("-- leave {ticket-id} | remove the car from the park \n")
	buffer.WriteString("-- search {color} | give a list of cars parked with that color \n")

	return buffer.String()
}
