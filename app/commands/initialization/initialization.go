package initialization

import (
	"github.com/parking-lot/app/client"
	"github.com/parking-lot/app/model/parking-lot"
	"strconv"
)

func Command(args []string, client *client.Client) {
	slots, err := strconv.Atoi(args[0])
	if err != nil {
		client.Send <- "Fail: The slots must be a number!"
		return
	}

	parkingLot := parking_lot.ParkingLot{}
	if canPark := parkingLot.SetSlots(slots); !canPark {
		client.Send <- "Parking lot already set"
		return
	}

	client.Parking = parkingLot
	client.Send <- "Parking lot generated " + args[0] + " slots"
}
