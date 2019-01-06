package search_test

import (
	"github.com/parking-lot/app/commands/park"
	"github.com/parking-lot/app/commands/search"
	"github.com/parking-lot/app/model/parking-lot"
	"strings"
	"testing"
)

func TestSearch(t *testing.T) {
	lot := parking_lot.ParkingLot{}
	_ = lot.SetSlots(3)

	ticket1, _ := park.Command(strings.Fields("HHH-TT red"), &lot)
	_ = lot.Park(*ticket1)

	ticket2, _ := park.Command(strings.Fields("HHH-TT blue"), &lot)
	_ = lot.Park(*ticket2)

	tickets, _ := search.Command(strings.Fields("red"), &lot)
	if len(tickets) != 1 {
		t.Error("Should have returned 1 car")
	}
	if strings.Compare(tickets[0].GetCar().Color, "red") != 0 {
		t.Error("Founded car is not red")
	}
}

func TestSearchNoColor(t *testing.T) {
	lot := parking_lot.ParkingLot{}
	_ = lot.SetSlots(3)

	ticket1, _ := park.Command(strings.Fields("HHH-TT red"), &lot)
	_ = lot.Park(*ticket1)

	if _, err := search.Command(strings.Fields(""), &lot); err == nil {
		t.Error("Wrong command should return error ")
	}
}
