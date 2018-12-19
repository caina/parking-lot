package parkingLot

import (
	"github.com/parking-lot/app/commands/initialization"
	"github.com/parking-lot/app/config"
	"testing"
)

func TestParkingLotCreation(t *testing.T) {

	mockClient := config.Client{
		Send: make(chan string),
		Stop: make(chan bool),
	}
	defer close(mockClient.Send)
	defer close(mockClient.Stop)

	args := []string{"39"}
	initialization.Command(args, &mockClient)

	if mockClient.Parking.Slots != 30 {
		t.Error("slots not created")
	}

}
