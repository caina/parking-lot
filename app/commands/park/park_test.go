package park_test

import (
	"github.com/parking-lot/app/commands/leave"
	"github.com/parking-lot/app/commands/park"
	"github.com/parking-lot/app/model/parking-lot"
	"strconv"
	"strings"
	"testing"
)

func TestCommandMalFormatted(t *testing.T) {
	lot := &parking_lot.ParkingLot{}
	_ = lot.SetSlots(1)
	if _, errors := park.Command(strings.Fields("MHT Red blue"), lot); errors != nil {
		t.Error("should give an error if its not only 2 commands")
	}
}

func TestPark(t *testing.T) {
	lot := &parking_lot.ParkingLot{}
	_ = lot.SetSlots(1)

	ticket, _ := park.Command(strings.Fields("HHH-TTT Red"), lot)
	if ticket.GetCar().Color != "Red" {
		t.Error("Ticket not properly sued")
	}
}

func TestParkingPosition(t *testing.T) {
	lot := &parking_lot.ParkingLot{}
	_ = lot.SetSlots(3)

	ticket1, _ := park.Command(strings.Fields("HHH-TTT Red"), lot)
	lot.Park(*ticket1)
	if ticket1.GetPosition() != 1 {
		t.Error("expected 1 got " + strconv.Itoa(ticket1.GetPosition()))
	}

	ticket2, _ := park.Command(strings.Fields("HHH-XXX Blue"), lot)
	lot.Park(*ticket2)
	if ticket2.GetPosition() != 2 {
		t.Error("expected 2 got " + strconv.Itoa(ticket2.GetPosition()))
	}

	ticket3, _ := park.Command(strings.Fields("HHH-ET Gray"), lot)
	lot.Park(*ticket3)
	if ticket3.GetPosition() != 3 {
		t.Error("expected 3 got " + strconv.Itoa(ticket3.GetPosition()))
	}
}

func TestOverlappingParkingPosition(t *testing.T) {
	lot := &parking_lot.ParkingLot{}
	_ = lot.SetSlots(3)

	firstCar, _ := park.Command(strings.Fields("HHH-TTT Red"), lot)
	_, _ = park.Command(strings.Fields("SSS-TTT Blue"), lot)
	_, _ = leave.Command(strings.Fields(firstCar.GetCode()), *lot)

	if ticketOnFirstPlace, _ := park.Command(strings.Fields("SSS-TTT Blue"), lot); ticketOnFirstPlace.GetPosition() != 1 {
		t.Error("Expected 1 got " + strconv.Itoa(ticketOnFirstPlace.GetPosition()))
	}
}

func TestParkCarWithLotNotInitialized(t *testing.T) {
	lot := &parking_lot.ParkingLot{}
	if _, err := park.Command(strings.Fields("HHH-TTT Red"), lot); err == nil {
		t.Error("Expected error for parking a car without a lot")
	}
}
