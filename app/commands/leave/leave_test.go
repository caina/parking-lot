package leave_test

import (
	"github.com/parking-lot/app/commands/leave"
	"github.com/parking-lot/app/commands/park"
	"github.com/parking-lot/app/commands/search"
	"github.com/parking-lot/app/model/parking-lot"
	"strconv"
	"strings"
	"testing"
)

func TestLeave(t *testing.T) {
	lot := parking_lot.ParkingLot{}
	_ = lot.SetSlots(3)

	ticket, _ := park.Command(strings.Fields("MHT-RR blue"), &lot)
	lot.Park(*ticket)

	tickets, _ := leave.Command(strings.Fields(ticket.GetCode()), lot)
	if len(tickets) != 0 {
		t.Error("Expected 0 cars, but got: " + strconv.Itoa(len(tickets)))
	}
}

func TestLeaveCarNotFound(t *testing.T) {
	lot := parking_lot.ParkingLot{}
	_ = lot.SetSlots(3)

	ticket, _ := park.Command(strings.Fields("MHT-RR blue"), &lot)
	lot.Park(*ticket)

	if _, err := search.Command(strings.Fields("NOT-A-PLATE"), &lot); err == nil {
		t.Error("should show an error")
	}
}
