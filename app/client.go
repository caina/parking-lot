package app

import . "github.com/parking-lot/app/model/parking-lot"

type Client struct {
	Send    chan string
	Stop    chan bool
	Parking ParkingLot
}
