package parking_lot

import (
	"bytes"
	. "github.com/parking-lot/app/model/ticket"
)

type ParkingLot struct {
	slots   int
	tickets []Ticket
}

func (parkingLot *ParkingLot) CanPark() bool {
	return parkingLot.slots > len(parkingLot.tickets)
}

func (parkingLot *ParkingLot) GetAvailableSlots() int {
	return parkingLot.slots - len(parkingLot.tickets)
}

func (parkingLot *ParkingLot) SetSlots(slots int) bool {
	if parkingLot.slots == 0 {
		parkingLot.slots = slots
		return true
	}
	return false
}

func (parkingLot *ParkingLot) Park(ticket Ticket) bool {
	if !parkingLot.CanPark() {
		return false
	}

	parkingLot.tickets = append(parkingLot.tickets, ticket)
	return true
}

func (parkingLot *ParkingLot) ToString() string {
	buffer := bytes.Buffer{}

	for i := 0; i < len(parkingLot.tickets); i++ {
		buffer.WriteString(parkingLot.tickets[i].ToString() + " \n")
	}

	return buffer.String()
}

func (parkingLot *ParkingLot) GetTickets() []Ticket {
	return parkingLot.tickets
}

func (parkingLot *ParkingLot) SetTickets(tickets []Ticket) {
	parkingLot.tickets = tickets
}
