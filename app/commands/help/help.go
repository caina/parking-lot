package help

import (
	"github.com/parking-lot/app/config"
)

func Command(teste *config.Client) {

	teste.Send <- "message"
	//fmt.Println("list of commands available")
	//fmt.Println("-- init")
	//fmt.Println("-- status | list of cars parked")
	//fmt.Println("-- park {plate color} | park the car on the slot nearby the door")
	//fmt.Println("-- leave {ticket-id} | remove the car from the park")
}
