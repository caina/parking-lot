package initialization

import (
	"github.com/parking-lot/app"
	"github.com/parking-lot/app/constants"
	"github.com/parking-lot/app/model/parking-lot"
	"strconv"
)

func Command(args []string, client *app.Client) {
	slots, err := strconv.Atoi(args[0])
	if err != nil {
		client.Send <- constants.SlotsMustBeNumber
		return
	}

	parkingLot := parking_lot.ParkingLot{}
	if err := parkingLot.SetSlots(slots); err != nil {
		client.Send <- err.Error()
		return
	}

	client.Parking = parkingLot
	client.Send <- "Parking lot generated " + args[0] + " slots"
}
